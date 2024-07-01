package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a password string and creates a hash for it
func HashPassword(password string)(string, error){
	if len(password) < 6 {
		return "", fmt.Errorf("len of password too short: needs at least len of %v", 6)
	}
	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(byteHash), nil
}

// Checks a password matches the stored hash in the database
func VerifyPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("passwords don't match: %v", err)
	}
	return nil
}