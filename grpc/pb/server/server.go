package main

import (
	"flag"
	pb "github.com/jkraj/jkraj-golang/grpc/pb"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"net"
	"google.golang.org/grpc"
)

type echoServer struct { }

func newEchoServer() pb.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	glog.Info(msg)
	return msg, nil
}

func Run() error {
	listen, err := net.Listen("tcp", "localhost:8999")
	if err != nil {
		glog.Error(err)
		return err
	}

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, newEchoServer())
	server.Serve(listen)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}