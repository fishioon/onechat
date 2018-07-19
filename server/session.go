package main

import (
	"math/rand"

	pb "github.com/fishioon/onechat/chat"
	//"math/rand"
)

const (
	msgbufSzie = 10
)

// Group ...
type Group struct {
	ID uint64
}

// AddSession ...
func (g *Group) AddSession(s *Session) {
}

// RemoveSession ...
func (g *Group) RemoveSession(s *Session) {
}

// Session ...
type Session struct {
	Sid    string //session id
	UID    string
	msgbuf chan *pb.Msg
	online bool
	//groups map[uint64]*Group
}

// NewSession ...
func NewSession(uid string) *Session {
	//sid := newSessionID()
	sid := RandStringBytes(32)
	return &Session{
		Sid:    sid,
		UID:    uid,
		msgbuf: make(chan *pb.Msg, msgbufSzie),
		//groups: make(map[uint64]*Group),
		online: true,
	}
}

func (s *Session) pubMsg(msg *pb.Msg) error {
	s.msgbuf <- msg
	return nil
}

func (s *Session) joinGroup(group *Group) {
	group.AddSession(s)
}

func (s *Session) leaveGroup(group *Group) {
	group.RemoveSession(s)
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
