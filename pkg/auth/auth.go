package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"database/sql"

	"github.com/Vladislavius12/Go_project/crypto"
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
	var pas string
	var name string
	var criptLogin string
	var criptPass string
	var descriptLogin string
	var descriptPass string

	name = req.GetLogin()
	pas = req.GetPassword()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	CheckError(err)
	criptLogin = crypto.Cripting(GetID(name), *privateKey)
	criptPass = crypto.Cripting(GetPassword(pas), *privateKey)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	descriptLogin = crypto.RSA_OAEP_Decrypt(criptLogin, *privateKey)
	descriptPass = crypto.RSA_OAEP_Decrypt(criptPass, *privateKey)

	if name == criptLogin && pas == criptPass {
		return &pb.AuthResponse{Res: "Hello"}, nil
	} else {
		fmt.Println(descriptLogin)
		fmt.Println(descriptPass)
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

func GetID(name string) string {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open db
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	row, err := db.Query("SELECT login FROM users where login='" + name + "'")
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
		fmt.Println(p.login)
		p.login = name
	}
	return name
}

func GetPassword(pas string) string {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open db
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	row, err := db.Query("SELECT pass FROM users where pass='" + pas + "'")
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
		fmt.Println(p.pass)
		p.pass = pas
	}
	return pas
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
