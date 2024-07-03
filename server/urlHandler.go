package server

import (
	"net/http"
	"strings"

	"errors"

	"github.com/ekefan/panda_url_shortner/authorize"
	"github.com/ekefan/panda_url_shortner/database"
	"github.com/ekefan/panda_url_shortner/util"
	"github.com/gin-gonic/gin"
)

// errorResponse converts the err to a string message and returns an ErrorResp struct
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
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
		// fmt.Println(req)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := ctx.MustGet(authPayloadKey)
	authPayload, ok := payload.(*authorize.Payload)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, 
			errorResponse(errors.New("not authorized")))
		return 
	}

	// generate short code
	shortCode, err := util.RandomShortCode(5)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//Args for creating a new URL in the database
	argsToSaveURL := database.CreateURLArgs{
		Owner: authPayload.Owner,
		ShortCode: shortCode,
		LongURL:   req.LongURL,
	}
	//Handle error properly, the error will not be unique constraint all the time
	savedURL, err := s.store.CreateURL(argsToSaveURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := struct{
		Owner string `json:"owner"`
		ShortCode string `json:"short_code"`
	}{
		Owner: savedURL.Owner,
		ShortCode: savedURL.ShortCode,
	}
	ctx.JSON(http.StatusOK, resp) //organise code
}

// GoToURLReq holds the uri value of the short code to redirect to
type GoToURLReq struct {
	ShortCode string `uri:"short_code" binding:"required"`
}

// goToURL redirects to longURL associated with the shortCode from the request
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

/// UpdateShortCodeReq holds field required for shortcode update
type UpdateShortCodeReq struct {
	ShortCode string `uri:"short_code" binding:"required"`
}

// UpdateUrlRequest holds the new short code to update
type UpdateUrlRequest struct {
	NewShortCode string `json:"new_short_code" binding:"required"`
}

// UrlResp holds fields required for the update response
type UrlResp struct {
	ShortCode string `json:"shortcode"`
	LongUrl   string `json:"long_url"`
}

// updateShortCode updates the short code of the user's URL
func (s *Server) updateShortCode(ctx *gin.Context) {
	var uriReq UpdateShortCodeReq
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var jsonReq UpdateUrlRequest
	if err := ctx.ShouldBindJSON(&jsonReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(authPayloadKey)
	authPayload, ok := payload.(*authorize.Payload)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errors.New("not authorized")))
		return
	}

	arg := database.TxUrlArgs{
		Owner:     authPayload.Owner,
		CurrentShortCode: uriReq.ShortCode,
		ShortCode: jsonReq.NewShortCode,
	}

	// Check if the new short code already exists for the owner
	existingUrl, err := s.store.GetUrlByOwnerAndShortCode(authPayload.Owner, jsonReq.NewShortCode)
	if err == nil && existingUrl.ShortCode != "" {
		ctx.JSON(http.StatusConflict, errorResponse(errors.New("short code already exists for this owner")))
		return
	}

	updatedUrl, err := s.store.TxUpdateShortCode(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := UrlResp{
		ShortCode: updatedUrl.ShortCode,
		LongUrl:   updatedUrl.LongURL,
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetUserUrlReq holds fields required for getting users urls
type GetUserUrlsReq struct {
	PageSize int `form:"page_size" binding:"required"`
	PageID int `form:"page_id" binding:"required"`
}

//getUserUrls server handler for getting user urls
func (s *Server) getUserUrls(ctx *gin.Context) {
	var req GetUserUrlsReq 
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := ctx.MustGet(authPayloadKey)
	authPayload, ok := payload.(*authorize.Payload)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, 
			errorResponse(errors.New("not authorized")))
		return 
	}
	arg := database.GetURLsArg{
		Owner: authPayload.Owner,
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}
	urls, err := s.store.GetURLs(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	urlResps := []UrlResp{}
	for _, url := range urls {
		urlResps = append(urlResps, UrlResp{
			ShortCode: url.ShortCode, 
			LongUrl: url.LongURL,
		})
	}
	ctx.JSON(http.StatusOK, urlResps)
}


func (s *Server) deleteUrl(ctx *gin.Context){
	var req UpdateShortCodeReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(authPayloadKey)
	authPayload, ok := payload.(*authorize.Payload)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, 
			errorResponse(errors.New("not authorized")))
		return 
	}
	arg := database.TxUrlArgs{
		Owner: authPayload.Owner,
		ShortCode: req.ShortCode,
	}

	//error if shortcode doesn't belong to user
	err := s.store.TxDeleteUrl(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := struct{ 
		Msg string `json:"msg"`
		Link string `json:"link"` 
		}{
			Msg: "successful",
			Link: ctx.Request.URL.String(),
		}
	ctx.JSON(http.StatusOK, resp)
}