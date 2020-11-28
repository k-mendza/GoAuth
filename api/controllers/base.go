package controllers

import (
	"GoAuth/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		log.Printf("cannot connect to %s database\n", DbDriver)
		log.Fatalf("error: %s\n", err)
	} else {
		log.Printf("sucessfuly connected to %s database", DbDriver)
	}

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Printf("server started on port %s\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
