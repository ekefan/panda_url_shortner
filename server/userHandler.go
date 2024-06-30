package server

import (
	"net/http"

	// "github.com/ekefan/panda_url_shortner/authorize"
	"github.com/ekefan/panda_url_shortner/database"
	"github.com/gin-gonic/gin"
)

// CreateUserReq holds necessary data for creating a new USER
type CreateUserReq struct {
	Name     string `json:"name" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResp struct{
		Name  string `json:"name"`
		Email string `json:"email"`
}
// createUser creates a new USER in database based on CreateUserReq fields
func (s *Server) createUser(ctx *gin.Context) {
	var req CreateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//implement password hash and verifier
	args := database.CreateUserArgs{
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: req.Password,
	}
	newUser, err := s.store.CreateUser(args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := UserResp{
		Name:  newUser.Name,
		Email: newUser.Email,
	}
	ctx.Redirect(http.StatusFound, "/user/login")
	ctx.JSON(http.StatusOK, resp)
}
//after creating user redirect to login user

type LoginRequest struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
func (s *Server) loginUser(ctx *gin.Context) {
	var req LoginRequest 
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//implement password hash and verifier
	user, err := s.store.GetUser(database.GetUserArgs{Name:req.Name})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	token, err := s.jwtMaker.CreateToken(req.Name, s.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := struct {
		Token string `json:"access_token"`
		Name string `json:"name"`
		Email string `json:"email"`

	}{
		Token : token,
		Name: user.Name,
		Email: user.Email,
	}
	ctx.JSON(http.StatusOK, resp)
}
