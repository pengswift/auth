package main

import (
	"log"

	pb "github.com/pengswift/auth/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:60010"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)

	r, err := c.Login(context.Background(), &pb.AuthRequest{
		Username: "路飞",
		Password: "橡皮人",
		SdkType:  pb.SDKType_SDK_OFFICIAL,
	})
	if err != nil {
		log.Fatalf("could not filter: %v", err)
	}
	log.Printf("Userid: %d\n", r.Userid)
	log.Printf("LastServerID: %d\n", r.LastServerID)
}
