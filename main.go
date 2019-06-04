package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/fishioon/onechat/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var Version, BuildTime string

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
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChatServer(grpcServer, NewChatServer())
	grpcServer.Serve(lis)
}
