package model

import "github.com/dgrijalva/jwt-go"

const TokenExpiration = 60 * 60 * 24 * 7 // 7 days

var passwordSgoalt = ""
var jwtSecret = ""

// hashes the password using Argon2id
func Argon(password string) {}

// generate a new JWT token for the user
func GenerateToken(username string) {}

// Verify Authorization
func VerifyAuth(auth string) {}

// Verify token
func VerifyToken(tokenUri string) {}

// Validate token
func ValidateToken(token *jwt.Token) {}
