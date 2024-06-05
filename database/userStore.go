package database

import (
)

// CreateUserArgs args for  creating a new user
type CreateUserArgs struct {
	Name string `json:"name"`
	Email string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// CreateUser database query function for creating a new user row
// on success it returns a USER object and a nil error
func (s *Query) CreateUser(args CreateUserArgs) (USER, error){
	userRow := USER{Name: args.Name, Email: args.Email, Password: args.HashedPassword}
	result := s.db.Create(&userRow)
	return userRow, result.Error
}

// GetUserArgs args for getting user from the database
type GetUserArgs struct {
	Name string `json:"name"`
}

// GetUser gets a USER object from the database
// on sucess it returns the USER object specified by the GetUserArgs
func (s *Query) GetUser(args GetUserArgs) (USER, error){
	userRow := USER{}
	result := s.db.Where("name = ?", args.Name).First(&userRow)
	return userRow, result.Error
}