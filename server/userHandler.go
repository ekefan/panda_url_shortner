package server

import "github.com/gin-gonic/gin"

//CreateUserReq holds necessary data for creating a new USER 
type CreateUserReq struct {
	Name string `json:"name" binding:"require, alphanum"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required, min=6"`
}
//createUser creates a new USER in database based on CreateUserReq fields
func (s *Server) createUser(ctx *gin.Context) {
	// var req CreateUserReq
}
func (s *Server) loginUser(ctx *gin.Context) {}



///after creating user redirect to login user