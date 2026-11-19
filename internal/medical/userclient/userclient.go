package userclient

import (
	"context"

	"github.com/teakingwang/backend/api/user"
	"github.com/teakingwang/backend/pkg/consul"
	"github.com/teakingwang/backend/pkg/logger"
	"google.golang.org/grpc"
)

type UserClient struct {
	discovery consul.Discovery
	conn      *grpc.ClientConn
	client    user.UserServiceClient
}

func NewUserClient(discovery consul.Discovery) (*UserClient, error) {
	conn, err := discovery.GetGRPCConn("user")
	if err != nil {
		return nil, err
	}
	return &UserClient{
		discovery: discovery,
		conn:      conn,
		client:    user.NewUserServiceClient(conn),
	}, nil
}

func (c *UserClient) GetUserNickname(ctx context.Context, userID int64) string {
	resp, err := c.client.GetUser(ctx, &user.GetUserRequest{Id: userID})
	if err != nil {
		logger.Warnf("get user nickname from user-service failed: %v", err)
		return ""
	}
	if resp.GetUsername() != "" {
		return resp.GetUsername()
	}
	return ""
}

func (c *UserClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}