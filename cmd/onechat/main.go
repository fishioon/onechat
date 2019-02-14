package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"strings"

	pb "github.com/fishioon/onechat/onechat"
	"github.com/fishioon/onechat/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

func serve(c *server.Config, logger *zap.Logger) error {
	grpcOptions := []grpc.ServerOption{grpc.UnaryInterceptor(ensureValidToken)}
	list, err := net.Listen("tcp", c.Addr)
	if err != nil {
		return fmt.Errorf("listening on %s failed: %v", c.Addr, err)
	}
	s := grpc.NewServer(grpcOptions...)
	chatServer := server.NewChatServer(c, logger)
	pb.RegisterChatServer(s, chatServer)

	logger.Info("server running, start listen", zap.String("addr", c.Addr))
	return s.Serve(list)
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8379", "onechat listen address")
	flag.Parse()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	c := &server.Config{
		Addr: *addr,
	}
	serve(c, logger)
}
