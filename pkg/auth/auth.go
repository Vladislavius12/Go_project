package auth

import (
	"context"

	pb "github.com/Vladislavius12/Go_project/pkg/api"
)

// AUTHServer ...
type AUTHServer struct{ pb.UnimplementedAuthServer }

// Auth ...
func (s *AUTHServer) Auth(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	var a string
	a = "123"
	var b string
	b = "123"
	if req.GetLogin() == a && req.GetPassword() == b {
		return &pb.AuthResponse{Res: "Hello"}, nil
	}
	return &pb.AuthResponse{Res: "Error"}, nil
}
