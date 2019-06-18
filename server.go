package main

import (
	"context"
	"log"
	"strings"

	pb "github.com/fishioon/onechat/chat"
	"github.com/go-redis/redis"
)

type Config struct {
	Addr string
}

type Stream struct {
	channel chan *pb.Msg
}

type Channel struct {
	streams map[string]*Stream
}

// ChatServer is used to implement chat
type ChatServer struct {
	rds          *redis.ClusterClient
	streams      map[string]*Stream
	channelNames []string
	channels     map[string]*Channel
}

// NewChatServer ...
func NewChatServer() *ChatServer {
	rds := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	rds.Ping()
	return &ChatServer{
		rds: rds,
	}
}

func (cs *ChatServer) Working() {
	res, err := cs.rds.XReadStreams(cs.channelNames...).Result()
	if err != nil {
		log.Printf("[redis] read streams err=%s", err.Error())
	}
	for _, e := range res {
		if e.Stream == "x:streams" {
		}
	}
}

func (cs *ChatServer) GetStream(token string) *Stream {
	s, ok := cs.streams[token]
	if !ok {
		s := &Stream{
			channel: make(chan *pb.Msg, 100),
		}
		cs.streams[token] = s
	}
	return s
}

// Conn ...
func (cs *ChatServer) Conn(in *pb.ConnReq, stream pb.Chat_ConnServer) error {
	s := cs.GetStream(in.GetToken())
	for {
		select {
		case msg := <-s.channel:
			if err := stream.Send(msg); err != nil {
				log.Printf("stream send fail %s", err.Error())
				return err
			}
		}
	}
	/*
		userAction := "x:action:" + uid
		streams := []string{userAction}
		for {
			res, err := cs.rds.XReadStreams(streams...).Result()
			for _, e := range res {
				if e.Stream == userAction {
					for _, msg := range e.Messages {
						if msg.Values["action"] == "group-action" {
						}
					}
				} else if strings.HasPrefix(e.Stream, "x:channels:") {
					for _, msg := range e.Messages {
						if err := stream.Send(msg); err != nil {
							log.Printf("stream send fail %s", err.Error())
							return err
						}
					}
				}
			}
		}
	*/
}

// Pub ...
func (cs *ChatServer) PubMsg(ctx context.Context, in *pb.PubMsgReq) (*pb.PubMsgRsp, error) {
	// s := getSession(ctx)
	return nil, nil
}

// Group ...
func (cs *ChatServer) GroupAction(ctx context.Context, req *pb.GroupActionReq) (resp *pb.GroupActionRsp, err error) {
	switch req.GetAction() {
	case "active":
	case "join":
	}
	return nil, nil
}

/*
func (cs *ChatServer) GetSessionByToken(token string) (*Session, error) {
	arr := strings.Split(token, "-")
}

func (cs *ChatServer) GetSession(uid, token string) (*Session, error) {
	ss, ok := cs.sessions[sid]
	if !ok {
		ss = NewSession(sid, uid)
		cs.sessions[sid] = ss
	}
	return ss, nil
}

func getSession(ctx context.Context) *Session {
	return ctx.Value("session").(*Session)
}
*/

func authToken(token string) string {
	return strings.Split(token, "-")[0]
}
