package controllers

import (
	"event_backend/api/middlewares"
	"event_backend/api/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // pg driver

	"github.com/qor/validations"
)

// Server => Server struct
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var errList = make(map[string]string)

// Initialize => Initialized DB connection
func (server *Server) Initialize(dbUser, dbPassword, dbPort, dbHost, dbName string) {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to postgres database")
		log.Fatal("This is the error connecting to postgres:", err)
	} else {
		fmt.Printf("We are connected to the postgres database")
	}

	validations.RegisterCallbacks(server.DB)

	// database migration
	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Profiles{},
		&models.Events{},
		&models.Post{},
	)

	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())

	server.intializeRoutes()

}

// Run => Runs server
func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
