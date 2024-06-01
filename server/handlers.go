package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) shortenURL(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"err": "no error"}) //organise code
}
func (s *Server) goToURL(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"err": "no error"}) //organise code
}

///Stopped at handling server request for creating new shortcode  and redirecting incomming one.
