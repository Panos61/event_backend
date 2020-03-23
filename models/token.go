package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//Token struct declaration
type Token struct {
	UserID   uint
	Username string
	Email    string
	*jwt.StandardClaims
}
