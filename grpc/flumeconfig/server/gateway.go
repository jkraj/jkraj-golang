package main

import (
	"flag"
	fc "github.com/jkraj/jkraj-golang/grpc/flumeconfig"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
	"github.com/golang/glog"
)

var (
	echoEndpoint = flag.String("flume_config_endpoint", "localhost:8999", "endpoint of FlumeConfig")
)

func RunEndPoint(address string, opts ...runtime.ServeMux) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := fc.RegisterFlumeConfigHandlerFromEndpoint(ctx, mux, *echoEndpoint, dialOpts)

	if err != nil {
		return err
	}

	http.ListenAndServe(address, mux)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := RunEndPoint(":8080"); err != nil {
		glog.Fatal(err)
	}
}
