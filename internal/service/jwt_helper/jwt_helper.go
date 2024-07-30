package jwt_helper

import (
	"crypto/rsa"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Config конфиг
type Config struct {
	PrivateKeyFile string `yaml:"private_key_file"`
	PublicKeyFile  string `yaml:"public_key_file"`
}

// Claims стандартные claims
type Claims interface {
	jwt.Claims
}

// JwtHelper jwt хелпер
type JwtHelper struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// NewHelper новый jwt хелпер
func NewHelper(cfg Config) (*JwtHelper, error) {
	buf, err := os.ReadFile(cfg.PrivateKeyFile)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(buf)
	if err != nil {
		return nil, err
	}

	buf, err = os.ReadFile(cfg.PublicKeyFile)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(buf)
	if err != nil {
		return nil, err
	}

	return &JwtHelper{
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

// Parse  validate, and return a token.
func (h *JwtHelper) Parse(token string, claims Claims) error {
	_, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
		return h.publicKey, nil
	})
	if err != nil {
		return err
	}
	return claims.Valid()
}

// CreateToken создание токена на основе claims
func (h *JwtHelper) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(h.privateKey)
}
