package session_manager

import (
	"github.com/dgrijalva/jwt-go"
)

// AccessSessionClaims данные авторизованного токена
type AccessSessionClaims struct {
	jwt.StandardClaims
	AccessSession
}

// Valid валидность claims
func (c *AccessSessionClaims) Valid() error {
	err := c.StandardClaims.Valid()
	if err != nil {
		return err
	}
	return nil
}

// RefreshSessionClaims данные токена обновления
type RefreshSessionClaims struct {
	jwt.StandardClaims
	RefreshSession
}

// Valid валидность claims
func (c *RefreshSessionClaims) Valid() error {
	err := c.StandardClaims.Valid()
	if err != nil {
		return err
	}

	return nil
}
