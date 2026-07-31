package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"

	"github.com/teakingwang/backend/api/user"
	"github.com/teakingwang/backend/config"
	"github.com/teakingwang/backend/internal/user/app"
	"github.com/teakingwang/backend/internal/user/controller"
	"github.com/teakingwang/backend/pkg/auth"
	"github.com/teakingwang/backend/pkg/consul"
	"github.com/teakingwang/backend/pkg/logger"
	"github.com/teakingwang/backend/pkg/utils/idgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	if err := logger.Init(true); err != nil {
		panic("logger init failed: " + err.Error())
	}

	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("panic occurred: %v", r)
			logger.Errorf("stack trace:\n%s", string(debug.Stack()))
		}
		logger.Sync()
	}()

	if err := config.LoadConfigFromConsul("config/user"); err != nil {
		logger.Warn("load from consul failed: %v", err)
		if err := config.LoadConfig(); err != nil {
			panic(fmt.Errorf("failed to load config: %v", err))
		}
	}

	auth.Init(config.Config.JWT.Secret, config.Config.JWT.ExpireSeconds)

	if err := idgen.Init(); err != nil {
		panic(fmt.Errorf("failed to initialize idgen: %v", err))
	}

	if err := run(); err != nil {
		logger.Errorf("service exited with error: %v", err)
	}

	waitForShutdown()
}

func run() error {
	ctx, err := app.NewAppContext()
	if err != nil {
		panic(fmt.Errorf("new appcontext err:%v", err))
	}

	go func() {
		httpAddr := ":" + config.Config.Server.User.HTTPPort
		router := controller.NewHTTPRouter(ctx.UserService, ctx.RBACService, ctx.MedicalRecordService)
		logger.Infof("HTTP server listening on %s", httpAddr)
		if err := http.ListenAndServe(httpAddr, router); err != nil {
			logger.Errorf("HTTP server error: %v", err)
		}
	}()

	go func() {
		grpcAddr := "0.0.0.0:" + config.Config.Server.User.GRPCPort
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Errorf("failed to listen grpc: %v", err)
			return
		}

		s := grpc.NewServer()
		registerHealthCheck(s)
		user.RegisterUserServiceServer(s, controller.NewUserController(ctx.UserService))

		consulClient, err := consul.NewConsulClient(config.Config.Consul.Address)
		if err != nil {
			logger.Errorf("failed to create consul client: %v", err)
			return
		}
		serviceID := config.Config.Server.User.Name + "-" + config.Config.Server.User.GRPCPort
		serviceName := config.Config.Server.User.Name
		serviceAddress := config.Config.Server.User.Host
		servicePort, err := strconv.Atoi(config.Config.Server.User.GRPCPort)
		if err != nil {
			logger.Errorf("invalid service port: %v", err)
			return
		}

		logger.Infof("Registering user service to Consul: %s", serviceID)
		if err := consulClient.RegisterService(serviceID, serviceName, serviceAddress, servicePort, []string{"grpc", "user"}); err != nil {
			logger.Errorf("consul register error: %v", err)
			return
		}
		defer consulClient.DeregisterService(serviceID)

		logger.Infof("gRPC server listening on %s", grpcAddr)
		if err := s.Serve(lis); err != nil {
			logger.Errorf("gRPC server failed: %v", err)
		}
	}()

	waitForShutdown()
	return nil
}

func waitForShutdown() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh
	logger.Info("Shutdown signal received, cleaning up...")
	logger.Sync()
	time.Sleep(time.Second)
	os.Exit(0)
}

func registerHealthCheck(s *grpc.Server) {
	hs := health.NewServer()
	hs.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	hs.SetServingStatus("user", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(s, hs)
}
