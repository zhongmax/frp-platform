package request

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        string
	ID          uint
	Username    string
	AuthorityId uint
}
