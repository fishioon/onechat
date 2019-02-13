package server

import (
	"math/rand"

	pb "github.com/fishioon/onechat/proto"
	//"math/rand"
)

const (
	msgbufSzie = 10
)

// Session ...
type Session struct {
	Sid    string //session id
	UID    string
	online bool
	msgch  chan *pb.Msg
	//groups map[uint64]*Group
}

// NewSession ...
func NewSession(sid, uid string) *Session {
	if sid == "" {
		sid = RandStringBytes(32)
	}
	return &Session{
		Sid:    sid,
		UID:    uid,
		online: true,
		msgch:  make(chan *pb.Msg, 10),
	}
}

func (s *Session) offline() {
	s.online = false
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes ...
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
