package main

import (
	"context"
	"crypto/tls"
	"flag"
	"io"
	"log"

	pb "github.com/fishioon/onechat/onechat"
	"golang.org/x/oauth2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func main() {
	serverAddr := flag.String("addr", "127.0.0.1:8379", "onechat server address")
	listen := flag.Bool("listen", false, "client listen msg")
	pubmsg := flag.String("pub", "", "client publish msg")
	flag.Parse()

	perRPC := oauth.NewOauthAccess(fetchToken())
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}),
		),
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatClient(conn)
	stream, err := c.Conn(context.TODO(), &pb.ConnReq{})
	if err != nil {
		log.Fatalf("connect fail err=%s", err.Error())
	}
	if *listen {
		go func() {
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
		}()
	}
	if *pubmsg != "" {
		c.Pub(context.TODO(), &pb.Msg{
			Content: *pubmsg,
		})
	}
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
