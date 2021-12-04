package auth

import (
	"context"
	"fmt"

	"database/sql"

	pb "github.com/Vladislavius12/Go_project/pkg/api"
	_ "github.com/lib/pq"
)

var lg, psw string
var id int

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "254312"
	dbname   = "serverusers"
)

type data struct {
	id    int
	login string
	pass  string
	data  int
}

// AUTHServer ...
type AUTHServer struct{ pb.UnimplementedAuthServer }

// Auth ...
func (s *AUTHServer) Auth(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	var choiceNull, newName, newPass string
	var a string
	a = "123"
	var b string
	b = "123"

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	row, err := db.Query("SELECT * FROM users where pass=('psw') and login=('lg')")
	if err != nil {
		panic(err)
	}
	defer row.Close()

	if req.GetLogin() == a && req.GetPassword() == b { //
		//row = nil
		return &pb.AuthResponse{Res: "Hello"}, nil
	} else {
		fmt.Println(lg)
		fmt.Println("user name not found. Would you like to register? (Y,N)")
		fmt.Scan(&choiceNull)
		switch choiceNull {
		case "Y":
			fmt.Println("Enter the desired name and password")
			fmt.Scan(&newName, &newPass)
			psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

			// open db
			db, err := sql.Open("postgres", psqlconn)
			CheckError(err)
			defer db.Close()

			_, err = db.Exec("insert into users (login, pass) values (" + newName + ", " + newPass + ")")
			if err != nil {
				panic(err)
			}

			break
		case "N":
			break
		default:
			fmt.Println("Please try again")
			break
		}
		return &pb.AuthResponse{Res: "Error"}, nil
	}
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

/*
func GetID() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open db
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	row, err := db.Query("SELECT login FROM users")
	if err != nil {
		panic(err)
	}
	defer row.Close()
	datas := []data{}

	for row.Next() {
		p := data{}
		err := row.Scan(&p.login)
		if err != nil {
			fmt.Println(err)
			continue
		}
		datas = append(datas, p)
	}
	for _, p := range datas {
		lg = p.login
		//fmt.Println(p.login)
	}
}

func GetPassword() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open db
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	row, err := db.Query("SELECT * FROM users where pass='psw' and login='lg'")
	if err != nil {
		panic(err)
	}
	defer row.Close()
	datas := []data{}

	for row.Next() {
		p := data{}
		err := row.Scan(&p.pass)
		if err != nil {
			fmt.Println(err)
			continue
		}
		datas = append(datas, p)
	}
	for _, p := range datas {
		psw = p.pass
		fmt.Println(row)
	}
}
*/
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
