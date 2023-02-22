package model

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/argon2"
)

const TokenExpiration = 60 * 60 * 24 * 7 // 7 days

var passwordSalt = []byte("")
var jwtSecret = ""
var N = 32768
var R = 8

// hashes the password using Argon2id
// N: 32678, r: 8, p: 1
// 32768: the N parameter, which controls the number of iterations used in the algorithm. In scrypt, N should be a power of 2, and this value is typically chosen to be as large as possible without making key generation too slow.
// 8: the r parameter, which controls the block size used in the algorithm. The block size should be a multiple of 8 and less than or equal to (2^32 - 1) / 128 / N“, in order to avoid memory issues.
// 1: the p parameter, which controls the parallelization factor used in the algorithm. The parallelization factor should be set to a value that is optimized for the available CPU and memory resources.
func Argon(password string) {
	passwordHash, err := argon2.IDKey([]byte(password), passwordSalt, 32768, 8, 1, MaxLength)
	if err != nil {
		return nil, err
	}
	return passwordHash, nil
}

// generate a new JWT token for the user
func GenerateToken(username string) {}

// Verify Authorization
func VerifyAuth(auth string) {}

// Verify token
func VerifyToken(tokenUri string) {}

// Validate token
func ValidateToken(token *jwt.Token) {}
