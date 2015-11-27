package main

import (
	"log"
	"net"

	pb "github.com/pengswift/auth/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":60010"
)

type server struct {
}

func (s *server) init() {

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ins := &server{}
	ins.init()

	//pb.
}
