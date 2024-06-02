package main

import (
	"log"

	"github.com/ekefan/panda_url_shortner/server"
	"github.com/gin-gonic/gin"
	// "time"
)

// main: the main entry point for the url_shorten server
func main() {
	gin.SetMode(gin.ReleaseMode)
	newServer := server.NewServer()

	port := "0.0.0.0:8080" //supposed to be environment variable
	newServer.SetupRouter()
	if err := newServer.StartServer(port); err!= nil {
		log.Fatal("could not start server: %v", err)
	}
	/*
		Couldn't separated migration process from the main code using gorm
		db.Migrator().DropTable(&database.URL{}) //so I used this till I got the
												// schema I wanted
	*/
	// Start Server



}

/*
	=============  TODO =================
	For user schema, add constraint ---Unique(user, shortCode) ---unique(usr, longurl)
	Add script for migrating the database down or updateURL
	Structure Project properly
	Use viper to load environment variables
*/