package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "github.com/fishioon/onechat/proto"
	"github.com/fishioon/onechat/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

/*
func commandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version and exit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`Onechat Version: %s
Build Time: %s
Go Version: %s
Go OS/ARCH: %s %s
`, version.Version, version.BuildTime, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		},
	}
}
*/

func serve(c *server.Config, logger *zap.Logger) error {
	testInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		logger.Info("metadata interceptor", zap.Any("md", md))
		return handler(ctx, req)
	}
	grpcOptions := []grpc.ServerOption{grpc.UnaryInterceptor(testInterceptor)}
	list, err := net.Listen("tcp", c.Addr)
	if err != nil {
		return fmt.Errorf("listening on %s failed: %v", c.Addr, err)
	}
	s := grpc.NewServer(grpcOptions...)
	chatServer := server.NewChatServer(c, logger)
	pb.RegisterChatServer(s, chatServer)
	err = s.Serve(list)
	return fmt.Errorf("listening on %s failed: %v", c.Addr, err)
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8379", "onechat listen address")
	flag.Parse()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	c := &server.Config{
		Addr: *addr,
	}
	serve(c, logger)
}
