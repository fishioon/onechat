package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strings"

	pb "github.com/fishioon/onechat/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var Version, BuildTime string

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func main() {
	address := flag.String("host", "127.0.0.1:9379", "onechat server listen address")
	flag.Parse()
	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChatServer(grpcServer, NewChatServer())
	grpcServer.Serve(lis)
}

func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	session, err := auth(ctx, info.Server.(*ChatServer))
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, "session", session)
	return handler(ctx, req)
}

func auth(ctx context.Context, cs *ChatServer) (*Session, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	authorization := md["authorization"]
	if len(authorization) < 1 {
		return nil, errInvalidToken
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	res := strings.Split(token, "-")
	return &Session{uid: res[0], sid: token}, nil
}
