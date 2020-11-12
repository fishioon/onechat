package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/fishioon/onechat/chat"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"
)

func main() {
	address := flag.String("host", "127.0.0.1:9981", "onechat server address")
	token := flag.String("token", "12345684234", "login token")
	flag.Parse()
	perRPC := oauth.NewOauthAccess(&oauth2.Token{
		AccessToken: *token,
	})
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
	}
	conn, err := grpc.Dial(*address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatClient(conn)

	go recvMsg(c, *token)
	go heartBeat(c, time.Second*10)
	readCommand(c)
	return
}

func heartBeat(c pb.ChatClient, d time.Duration) {
	ticker := time.NewTicker(d)
	for _ = range ticker.C {
		c.HeartBeat(context.TODO(), nil)
	}
}

func recvMsg(c pb.ChatClient, token string) error {
	// connect to onechat
	connReq := &pb.ConnReq{
		Token: token,
	}
	stream, err := c.Conn(context.TODO(), connReq)
	if err != nil {
		log.Fatalf("conn fail err=%s", err.Error())
	}
	msg := &pb.Msg{}
	for {
		if err = stream.RecvMsg(msg); err != nil {
			log.Fatalf("stream recvmsg fail err=%s", err.Error())
			return err
		}
		log.Printf("recv msg: %+v", msg)
	}
}

func readCommand(c pb.ChatClient) (err error) {
	scanner := bufio.NewScanner(os.Stdin)
	packMsg := func(text string) *pb.PubMsgReq {
		data := strings.SplitN(text, " ", 2)
		return &pb.PubMsgReq{
			Msg: &pb.Msg{
				FromId:  "ifish",
				ToId:    data[0],
				Content: data[1],
			},
		}
	}
	for scanner.Scan() {
		msg := packMsg(scanner.Text())
		action := &pb.GroupActionReq{
			Gid:    msg.GetMsg().GetToId(),
			Action: "join",
		}
		if _, err = c.GroupAction(context.TODO(), action); err != nil {
			log.Printf("group action err=%s", err.Error())
			return
		}
		if _, err = c.PubMsg(context.TODO(), msg); err != nil {
			log.Printf("pub msg err=%s", err.Error())
			return
		}
	}
	return
}
