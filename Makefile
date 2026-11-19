# 项目模块名（用于生成 Go 代码中的 import path）
MODULE=github.com/teakingwang/ginmicro

# proto 根目录
PROTO_DIR=api

# 生成目录（一般与 proto 保持一致）
OUT_DIR=api

# 所有 proto 文件
PROTO_FILES=$(shell find $(PROTO_DIR) -name "*.proto")

# google api 文件路径（如 annotations.proto 和 http.proto）
GOOGLE_API_DIR=third_party

# protoc 编译选项
PROTOC_GEN_GO=protoc \
	--proto_path=$(PROTO_DIR) \
	--proto_path=$(GOOGLE_API_DIR) \
	--go_out=$(OUT_DIR) \
	--go-grpc_out=$(OUT_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$(OUT_DIR) \
	--grpc-gateway_opt=paths=source_relative,logtostderr=true

# 默认目标：编译 proto 文件
.PHONY: proto
proto:
	@echo "Generating gRPC code from proto files..."
	@echo $(PROTOC_GEN_GO) $(PROTO_FILES)
	@$(PROTOC_GEN_GO) $(PROTO_FILES)
	@echo "✅ Done."

# 构建所有微服务
.PHONY: build
build:
	@echo "🔨 Building all services..."
	go build -o bin/user-service ./cmd/user
	go build -o bin/medical-service ./cmd/medical
	@echo "✅ Build completed."

# 清理生成的二进制文件
.PHONY: clean
clean:
	@echo "🧹 Cleaning up..."
	rm -rf bin/*
	find . -name "*.pb.go" -delete
	find . -name "*.gw.go" -delete
	@echo "✅ Clean done."

# 运行 user-service
.PHONY: run-user
run-user:
	go run ./cmd/user

# 运行 medical-service
.PHONY: run-medical
run-medical:
	go run ./cmd/medical