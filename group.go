package main

import (
	pb "github.com/fishioon/onechat/chat"
)

// Group ...
type Group struct {
	ID       uint64
	sessions map[string]*Session
}

// AddSession ...
func (g *Group) AddSession(s *Session) {
}

// RemoveSession ...
func (g *Group) RemoveSession(s *Session) {
}

// Pub ...
func (g *Group) Pub(msg *pb.Msg) error {
	return nil
}
