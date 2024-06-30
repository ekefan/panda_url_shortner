package server

import (
	"fmt"
	"log"

	"github.com/ekefan/panda_url_shortner/authorize"
	"github.com/ekefan/panda_url_shortner/database"
	"github.com/ekefan/panda_url_shortner/util"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	store database.Store
	router *gin.Engine
	config util.Config
	jwtMaker *authorize.JwtMaker
}

func NewServer(config util.Config) *Server{
	dbConn, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting with database", err)
	}
	s := database.NewStore(dbConn)
	err = s.RunMigrations(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	makeJwt, err := authorize.NewJwtToken(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("can not create tokens: ", err)
	}
	return &Server{
		store: s,
		config: config,
		jwtMaker: makeJwt,
	}
}


func (s *Server) StartServer() error {
	if err := s.router.Run(); err != nil {
		return fmt.Errorf("error starting server: %s", err)
	}
	return nil
}


func (s *Server) SetupRouter() {
	newRouter := gin.Default()

	newRouter.POST("/user", s.createUser)
	newRouter.POST("/user/login", s.loginUser)

	newRouter.POST("/new", s.shortenURL)
	newRouter.GET("/:short_code", s.goToURL)
	
	s.router = newRouter
}


