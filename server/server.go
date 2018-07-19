package main

import (
	"context"

	pb "github.com/fishioon/onechat/chat"
)

// ChatServer is used to implement chat
type ChatServer struct {
	sessions map[string]*Session
}

// Conn ...
func (cs *ChatServer) Conn(in *pb.ConnReq, stream pb.Chat_ConnServer) error {
	s := getSession(stream.Context())
	defer s.offline()
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case msg := <-s.msgbuf:
			if err := stream.Send(msg); err != nil {
				return err
			}
		}
	}
}

// Pub ...
func (cs *ChatServer) Pub(ctx context.Context, in *pb.PubReq) (*pb.PubRsp, error) {
	s := getSession(ctx)
	if err := s.pubMsg(in.Msg); err != nil {
		return nil, err
	}
	return &pb.PubRsp{}, nil
}

// Group ...
func (cs *ChatServer) Group(ctx context.Context, in *pb.GroupReq) (*pb.GroupRsp, error) {
	return nil, nil
}

func getSession(ctx context.Context) *Session {
	return ctx.Value("session").(*Session)
}
