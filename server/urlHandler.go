package server

import (
	"fmt"
	"net/http"
	"strings"

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
	ShortCode string `uri:"short_code" binding:"required"`
}

func (s *Server) goToURL(ctx *gin.Context) {
	var req GoToURLReq
	//bind the uri to get the shortCode from the request uri
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//make call to the database to retrieve long url
	arg := database.GetURLArgs{
		ShortCode: req.ShortCode,
	}
	dbURL, err := s.store.GetURL(arg)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, dbURL.LongURL)
}