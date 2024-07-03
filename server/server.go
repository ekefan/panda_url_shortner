package server

import (
	"fmt"
	"log"

	"time"
	"github.com/ekefan/panda_url_shortner/authorize"
	"github.com/ekefan/panda_url_shortner/database"
	"github.com/ekefan/panda_url_shortner/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	store    database.Store
	router   *gin.Engine
	config   util.Config
	jwtMaker *authorize.JwtMaker
}

func NewServer(config util.Config) *Server {
	dbConn, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting with database", err)
	}
	s := database.NewStore(dbConn)
	err = s.RunMigrations(dbConn, 0)
	if err != nil {
		log.Fatal(err)
	}
	makeJwt, err := authorize.NewJwtToken(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("can not create tokens: ", err)
	}
	return &Server{
		store:    s,
		config:   config,
		jwtMaker: makeJwt,
	}
}

func (s *Server) StartServer() error {
	if err := s.router.Run(s.config.ServerAddress); err != nil {
		return fmt.Errorf("error starting server: %s", err)
	}
	return nil
}

func (s *Server) SetupRouter() {
	newRouter := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5500"}, // Change to your frontend's origin
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	newRouter.Use(cors.New(corsConfig))
	newRouter.POST("/user", s.createUser)
	newRouter.POST("/user/login", s.loginUser)

	routes := newRouter.Use(jwtAuthHandler(*s.jwtMaker))
	routes.POST("/new", s.shortenURL)
	routes.GET("/:short_code", s.goToURL)
	routes.GET("/user/login", s.getUserUrls)
	routes.PUT("/user/url/:short_code", s.updateShortCode)
	routes.DELETE("/user/url/:short_code", s.deleteUrl)

	// ================= URLS ================
	// 			getUrls --- query..with page size and offset //working on it // done
	// 			deleteUrl --- delete from url where id = url = ID //working on it // done
	// 			editshortcode -- transaction, begin,  udpate url. commit ///working on it // done

	// ================= USERs ===============

	///Working on these now.......
	//			updateUser ---transaction // begin update url.owner update user.name commit if any error rollback
	//			deleteUser --- transaction // begin delete urls where owner = username, delete users where name = username if any error rollback

	s.router = newRouter
}
