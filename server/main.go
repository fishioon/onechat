package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/fishioon/onechat/chat"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 8379, "chat listen port")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chatSrv := new(ChatServer)
	s := grpc.NewServer()
	pb.RegisterChatServer(s, chatSrv)
	s.Serve(lis)
}
