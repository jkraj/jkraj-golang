package main

import (
	"flag"
	fc "github.com/jkraj/jkraj-golang/grpc/flumeconfig"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"net"
	"google.golang.org/grpc"
)

type flumeConfig struct { }

func newFlumeConfig() fc.FlumeConfigServer {
	return new(flumeConfig)
}

func (s *flumeConfig) GetConfig(ctx context.Context, cfg *fc.GetCfg) (*fc.Config, error) {
	glog.Info("test")
	return nil, nil
}

func (s *flumeConfig) CreateConfig(ctx context.Context, cfg *fc.Config) (*fc.Config, error)  {
	return cfg, nil
}

func Run() error {
	listen, err := net.Listen("tcp", "localhost:8999")
	if err != nil {
		glog.Error(err)
		return err
	}
	server := grpc.NewServer()
	fc.RegisterFlumeConfigServer(server, newFlumeConfig())
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