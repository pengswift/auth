package main

import (
	"fmt"

	pb "github.com/pengswift/auth/auth"
)

func (s *server) loginByOfficial(name string, password string) (*pb.AuthResponse, error) {
	if name == "" || password == "" {
		return nil, fmt.Errorf("name or password is null")
	}

	if name == "路飞" && password == "橡皮人" {
		return &pb.AuthResponse{
			Userid:       10,
			LastServerID: -1,
		}, nil
	}
	return nil, fmt.Errorf("login failed")
}
