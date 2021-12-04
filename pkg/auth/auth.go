package auth

import (
	"context"
	"fmt"

	"database/sql"

	pb "github.com/Vladislavius12/Go_project/pkg/api"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "254312"
	dbname   = "serverusers"
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

func Connect() {
	// connect
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open db
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
