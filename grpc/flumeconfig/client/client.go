package main

import (
	"log"
	"google.golang.org/grpc"
	fc "github.com/jkraj/jkraj-golang/grpc/flumeconfig"

	"golang.org/x/net/context"
)

const (
	address = "localhost:8999"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fc.NewFlumeConfigClient(conn)

	// Contact the server and print out its response.

	cvt := fc.Converter{Type: "test-covt", DnsViewDefault: "Internal"}
	incp := fc.Interceptors{Name: []string{"converter"}, Converter: []*fc.Converter{&cvt}}
	ss := fc.SyslogSource{Type: "test-syslog"}
	wapi := fc.Wapi{Type: "test-wapi"}
	grid := fc.Grid{Type: "test-grid", SpoolDir: "/var/spool", Channels: []string{"ch1", "ch2"}, Interceptors: &incp}
	src := fc.Sources{Grid: &grid, Wapi: &wapi, Ss: &ss}
	agent := fc.Agent{Name: "agent1", Sources: &src}

	r, err := c.CreateConfig(context.Background(), &fc.Config{Agent: []*fc.Agent{&agent}})
	if err != nil {
		log.Fatalf("could create config: %v", err)
	}
	log.Printf("config value: %s", r.GetAgent())
}
