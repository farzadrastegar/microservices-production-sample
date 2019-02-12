//original video: https://www.youtube.com/watch?v=bM6N-vgPlyQ

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"microservices/homepage"
	"log"
	"microservices/server"
	"net/http"
	"os"
)

func init() {
	os.Setenv("MicroCertFile", "sslfiles/cert.pem")
	os.Setenv("MicroKeyFile", "sslfiles/key.pem")
	os.Setenv("MicroServiceAddr", ":8081")
}

func main() {
	logger := log.New(os.Stdout, "MyLogger ", log.LstdFlags | log.Lshortfile)

	//DB connection
	//db := NewDBConnection(logger)

	h := homepage.NewHandlers(logger, nil) //db)

	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	//the configuration below is to make a secure webserver
	//URL> https://blog.cloudflare.com/exposing-go-on-the-internet/
	srv := server.New(mux, os.Getenv("MicroServiceAddr"))

	logger.Println("server starting")

	err := srv.ListenAndServeTLS(os.Getenv("MicroCertFile"), os.Getenv("MicroKeyFile"))
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

func NewDBConnection(logger *log.Logger) (*sqlx.DB) {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:8889)/test")
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	//testing a DB query
	databases := []string{}
	db.Select(&databases, "show databases")
	logger.Println("databases in mysql are...")
	logger.Println(databases)

	return db
}

