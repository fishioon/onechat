package main

import (
	"context"
	"errors"
	"log"

	"github.com/bwmarrin/snowflake"
	pb "github.com/fishioon/onechat/chat"
)

type Config struct {
	Addr string
}

type Stream struct {
	token   string
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
	snow    *snowflake.Node
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
	snow, _ := snowflake.NewNode(1)
	cs := &ChatServer{
		snow:    snow,
		in:      make(chan *pb.Msg, 1000),
		action:  make(chan Action, 1000),
		streams: make(map[string]*Stream),
		groups:  make(map[string]*Group),
	}
	go cs.Working()
	return cs
}

func (cs *ChatServer) Working() {
	for {
		select {
		case msg := <-cs.in:
			log.Printf("recv msg: [%v]", msg)
			if group, ok := cs.groups[msg.GetToId()]; ok {
				for _, e := range group.streams {
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
		s = &Stream{
			token:   token,
			channel: make(chan *pb.Msg, 100),
		}
		cs.streams[token] = s
	}
	return s
}

// Conn ...
func (cs *ChatServer) Conn(in *pb.ConnReq, stream pb.Chat_ConnServer) error {
	token := in.GetToken()
	if token == "" {
		return errors.New("token empty")
	}
	s := cs.GetStream(token)
	defer func() {
		s.online = false
	}()
	s.online = true
	log.Printf("stream connect success: [%+v]", s)
	for msg := range s.channel {
		log.Printf("send msg [%s] to stream [%s]", msg.GetToId(), in.GetToken())
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
	msg.MsgId = cs.snow.Generate().String()
	err := cs.sendMsgToChannel(msg)
	return &pb.PubMsgRsp{}, err
}

func (cs *ChatServer) HeartBeat(ctx context.Context, req *pb.HeartBeatReq) (*pb.HeartBeatRsp, error) {
	return &pb.HeartBeatRsp{}, nil
}

// Group ...
func (cs *ChatServer) GroupAction(ctx context.Context, req *pb.GroupActionReq) (resp *pb.GroupActionRsp, err error) {
	ses := ctx.Value("session").(*Session)
	switch req.GetAction() {
	case "active":
	case "join":
		group := cs.GetGroup(req.GetGid())
		if s, ok := group.streams[ses.sid]; ok {
			s.online = true
		} else {
			stream, ok := cs.streams[ses.sid]
			if !ok {
				return nil, errors.New("need connect first")
			}
			group.streams[ses.sid] = stream
		}
		log.Printf("user [%s %s] join [%s]", ses.uid, ses.sid, group.id)
	}
	return &pb.GroupActionRsp{}, nil
}

func (cs *ChatServer) GetGroup(gid string) *Group {
	group, ok := cs.groups[gid]
	if !ok {
		group = &Group{
			id:      gid,
			streams: make(map[string]*Stream),
		}
		cs.groups[gid] = group
	}
	return group
}
