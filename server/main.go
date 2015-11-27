package main

import (
	"fmt"
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

func (s *server) Login(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	if in.SdkType == pb.SDKType_SDK_OFFICIAL {
		return s.loginByOfficial(in.Username, in.Password)
	}
	//else if in.SdkType == pb.SDKType_SDK_QQ {
	//	return s.loginByQQ(in.Username, in.Password)
	//} else if in.SdkType == pb.SDKType_SDK_WEIXIN {
	//	return s.loginByWeiXin(in.Username, in.Password)
	//} else if in.SdkType == pb.SDKType_SDK_WEIBO {
	//	return s.loginByWeiBo(in.Username, in.Password)
	//}

	return nil, fmt.Errorf("invalid sdktype %d", in.SdkType)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ins := &server{}
	ins.init()

	pb.RegisterAuthServiceServer(s, ins)

	s.Serve(lis)
}
