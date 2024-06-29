package server

import (
	"net/http"

	"github.com/ekefan/panda_url_shortner/database"
	"github.com/gin-gonic/gin"
)

// CreateUserReq holds necessary data for creating a new USER
type CreateUserReq struct {
	Name     string `json:"name" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// createUser creates a new USER in database based on CreateUserReq fields
func (s *Server) createUser(ctx *gin.Context) {
	var req CreateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
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
	resp := struct {
		Name  string `json: "name"`
		Email string `json:"email"`
	}{
		Name:  newUser.Name,
		Email: newUser.Email,
	}
	ctx.JSON(http.StatusOK, resp)
}
func (s *Server) loginUser(ctx *gin.Context) {}

///after creating user redirect to login user
