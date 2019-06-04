package main

import (
	"context"
	"errors"

	pb "github.com/fishioon/onechat/chat"
)

type Config struct {
	Addr string
}

// ChatServer is used to implement chat
type ChatServer struct {
	sessions map[string]*Session
	groups   map[string]*Group
}

// NewChatServer ...
func NewChatServer() *ChatServer {
	return &ChatServer{
		sessions: make(map[string]*Session),
		groups:   make(map[string]*Group),
	}
}

// GetGroup ...
func (cs *ChatServer) GetGroup(gid string) (*Group, error) {
	if group, ok := cs.groups[gid]; ok {
		return group, nil
	}
	return nil, errors.New("invalid group id")
}

// Conn ...
func (cs *ChatServer) Conn(in *pb.ConnReq, stream pb.Chat_ConnServer) error {
	s := getSession(stream.Context())
	defer s.offline()
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case msg := <-s.msgch:
			if err := stream.Send(msg); err != nil {
				return err
			}
		}
	}
}

// Pub ...
func (cs *ChatServer) Pub(ctx context.Context, in *pb.PubReq) (*pb.PubRsp, error) {
	// s := getSession(ctx)
	msg := in.GetMsg()
	if msg.GetMsgType() == pb.MsgType_GROUP {
		group, err := cs.GetGroup(msg.GetToId())
		if err != nil {
			return nil, err
		}
		if err = group.Pub(msg); err != nil {
			return nil, err
		}
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
