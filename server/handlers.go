package server

import (
	"fmt"
	"net/http"

	"github.com/ekefan/panda_url_shortner/database"
	"github.com/ekefan/panda_url_shortner/util"
	"github.com/gin-gonic/gin"
)

// ErrorResp custom error response for handler functions to be JSONified by ctx.JSON
type ErrorResp struct {
	err string
}

// errorResponse converts the err to a string message and returns an ErrorResp struct
func errorResponse(err error) ErrorResp {
	return ErrorResp{
		err: err.Error(),
	}
}

// shortenURLReq respresents http.request body for shortenURL handler
type ShortenURLReq struct {
	LongURL string `json:"long_url" binding:"required"`
}
// shortenURL handler creates a shortened url, returns the short-url in the resp body
func (s *Server) shortenURL(ctx *gin.Context) {
	var req ShortenURLReq
	//bind request to get LongURL
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(req)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// generate short code
	shortCode, err := util.RandomShortCode()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} 
	//Args for creating a new URL in the database
	argsToSaveURL := database.CreateURLArgs{
		ShortCode: shortCode,
		LongURL: req.LongURL,
	}
	//Handle error properly, the error will not be unique constraint all the time
	savedURL, err := s.store.CreateURL(argsToSaveURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("shortCode already exists: %v", err)))
		return
	}
	ctx.JSON(http.StatusOK, savedURL.ShortCode) //organise code
}


type GoToURLReq struct {
	ShortCode string `json:"short_code" binding:"required"`
}

func (s *Server) goToURL(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"err": "no error"})
}


///Stopped at handling server request for creating new shortcode  and redirecting incomming one.