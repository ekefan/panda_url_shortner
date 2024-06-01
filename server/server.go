package server

import (
	"fmt"
	"log"

	"github.com/ekefan/panda_url_shortner/database"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	store database.Store
	router *gin.Engine
}

func NewServer() *Server{
	dbConn, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting with database", err)
	}
	s := database.NewStore(dbConn)
	err = s.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		store: s,
	}
}


func (s *Server) StartServer(port string) error {
	if err := s.router.Run(); err != nil {
		return fmt.Errorf("error starting server: %s", err)
	}
	return nil
}


func (s *Server) SetupRouter() {
	router := gin.Default()

	router.POST("/new", s.shortenURL)
	router.GET("/:short_code", s.goToURL)

}