package main

import (
	"context"
	"errors"
	"log"
	"strings"

	pb "github.com/fishioon/onechat/chat"
)

type Config struct {
	Addr string
}

type Stream struct {
	channel chan *pb.Msg
	online  bool
}

type Session struct {
	sid string
	uid string
}

type Group struct {
	id      string
	uri     string
	streams map[string]*Stream
}

// ChatServer is used to implement chat
type ChatServer struct {
	streams map[string]*Stream
	groups  map[string]*Group
	in      chan *pb.Msg
	action  chan Action
}

type Action struct {
	action string
	id     string
}

// NewChatServer ...
func NewChatServer() *ChatServer {
	return &ChatServer{
		in:      make(chan *pb.Msg, 1000),
		action:  make(chan Action, 1000),
		streams: make(map[string]*Stream),
		groups:  make(map[string]*Group),
	}
}

func (cs *ChatServer) Working() {
	for {
		select {
		case msg := <-cs.in:
			if ch, ok := cs.groups[msg.GetToId()]; ok {
				for _, e := range ch.streams {
					if e.online {
						e.channel <- msg
					}
				}
			}
		case action := <-cs.action:
			switch action.action {
			case "leave":
				delete(cs.groups, action.id)
			case "dead":
				delete(cs.streams, action.id)
			case "join":
			}
		}
	}
}

func (cs *ChatServer) sendMsgToChannel(msg *pb.Msg) error {
	cs.in <- msg
	return nil
}

func (cs *ChatServer) GetStream(token string) *Stream {
	s, ok := cs.streams[token]
	if !ok {
		s := &Stream{
			channel: make(chan *pb.Msg, 100),
			online:  true,
		}
		cs.streams[token] = s
	}
	return s
}

// Conn ...
func (cs *ChatServer) Conn(in *pb.ConnReq, stream pb.Chat_ConnServer) error {
	s := cs.GetStream(in.GetToken())
	defer func() {
		s.online = false
	}()
	for msg := range s.channel {
		if err := stream.Send(msg); err != nil {
			log.Printf("stream send fail %s", err.Error())
			return err
		}
	}
	return nil
}

// Pub ...
func (cs *ChatServer) PubMsg(ctx context.Context, req *pb.PubMsgReq) (*pb.PubMsgRsp, error) {
	msg := req.GetMsg()
	err := cs.sendMsgToChannel(msg)
	return &pb.PubMsgRsp{}, err
}

// Group ...
func (cs *ChatServer) GroupAction(ctx context.Context, req *pb.GroupActionReq) (resp *pb.GroupActionRsp, err error) {
	s := ctx.Value("session").(*Session)
	switch req.GetAction() {
	case "active":
	case "join":
		group := cs.GetGroup(req.GetGid())
		if _, ok := group.streams[s.sid]; ok {
			return nil, nil
		}
		stream, ok := cs.streams[s.sid]
		if !ok {
			return nil, errors.New("need connect first")
		}
		group.streams[s.sid] = stream
	}
	return nil, nil
}

func (cs *ChatServer) GetGroup(gid string) *Group {
	return nil
}

func authToken(token string) string {
	return strings.Split(token, "-")[0]
}
