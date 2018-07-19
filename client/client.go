package main

import (
	"context"
	"io"
	"log"

	pb "github.com/fishioon/onechat/chat"
	"golang.org/x/oauth2"

	"google.golang.org/grpc"
)

func main() {
	/*
		perRPC := oauth.NewOauthAccess(fetchToken())
		opts := []grpc.DialOption{
			grpc.WithPerRPCCredentials(perRPC),
			grpc.WithTransportCredentials(
				credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}),
			),
		}
	*/
	conn, err := grpc.Dial(":8080")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatClient(conn)
	stream, err := c.Conn(context.TODO(), &pb.ConnReq{})
	if err != nil {
		log.Fatalf("connect fail err=%s", err.Error())
	}
	for {
		feature, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.Conn(_) = _, %v", c, err)
		}
		log.Println(feature)
	}
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "chat-secret",
	}
}
