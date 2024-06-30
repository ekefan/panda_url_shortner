package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ekefan/panda_url_shortner/authorize"
	"github.com/gin-gonic/gin"
)

// header variables that must exist
const (
	authHeaderKey  string = "authorization"
	authType       string = "bearer"
	authPayloadKey string = "authorization_payload"
)

func jwtAuthHandler(token authorize.JwtMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authHeaderKey)
		//check that the authHeader is provided
		if len(authHeader) == 0 {
			err := errors.New("authorisation header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		//split the authHeader into fields
		authFields := strings.Fields(authHeader)
		//check the first field and make sure it's a bearer
		if len(authFields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		if authType != strings.ToLower(authFields[0]) {
			err := fmt.Errorf("invalid authorization type %v", authFields[0])
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		//get the token from the second string
		accessToken := authFields[1]
		//verify the token
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		//set a new header, authpayloadkey and set it's value to payload
		ctx.Set(authPayloadKey, payload)
		ctx.Next()

	}
}
