package main

import (
	"context"

	pb "github.com/fishioon/onechat/chat"
)

// ChatServer is used to implement chat
type ChatServer struct{}

// Conn ...
func (s *ChatServer) Conn(in *pb.ConnReq, stream pb.Chat_ConnServer) error {
	return nil
}

// Pub ...
func (s *ChatServer) Pub(ctx context.Context, in *pb.PubReq) (*pb.PubRsp, error) {
	return nil, nil
}

// Group ...
func (s *ChatServer) Group(ctx context.Context, in *pb.GroupReq) (*pb.GroupRsp, error) {
	return nil, nil
}

func getUID(ctx context.Context) string {
	return ctx.Value("uid").(string)
}
