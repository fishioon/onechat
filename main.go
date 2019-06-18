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
	"google.golang.org/grpc/credentials"
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
	tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile := flag.String("cert_file", "", "The TLS cert file")
	keyFile := flag.String("key_file", "", "The TLS key file")

	flag.Parse()
	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
	}
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChatServer(grpcServer, NewChatServer())
	grpcServer.Serve(lis)
}

func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := auth(ctx, info.Server.(*ChatServer)); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func auth(ctx context.Context, cs *ChatServer) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errMissingMetadata
	}
	authorization := md["authorization"]
	if len(authorization) < 1 {
		return errInvalidToken
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	res := strings.Split(token, "-")
	ctx = context.WithValue(ctx, "uid", res[0])
	ctx = context.WithValue(ctx, "accesstoken", res[1])
	return nil
}
