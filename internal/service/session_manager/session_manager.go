package session_manager

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/server"
	"github.com/photo-pixels/platform/serviceerr"

	"github.com/photo-pixels/user-account/internal/service/jwt_helper"
)

var (
	// ErrSessionNotFound сессия авторизованного пользователя не найдена
	ErrSessionNotFound = errors.New("session not found")
)

// Config конфигурация выпуска jwt
type Config struct {
	Audience             string        `yaml:"audience"`
	Issuer               string        `yaml:"issuer"`
	AccessTokenDuration  time.Duration `yaml:"access_token_duration"`
	RefreshTokenDuration time.Duration `yaml:"refresh_token_duration"`
}

// JWTHelper хелпер для работы с jwt
type JWTHelper interface {
	Parse(token string, claims jwt_helper.Claims) error
	CreateToken(claims jwt_helper.Claims) (string, error)
}

// SessionManager менеджер работы с токенами и данными авторизованного пользователя
type SessionManager struct {
	logger    log.Logger
	cfg       Config
	jwtHelper JWTHelper
}

// NewSessionManager новый менеджер работы с токенам
func NewSessionManager(
	logger log.Logger,
	cfg Config,
	jwtHelper JWTHelper,
) *SessionManager {
	return &SessionManager{
		logger:    logger.Named("session_manager"),
		cfg:       cfg,
		jwtHelper: jwtHelper,
	}
}

// CreateTokenByAccessSession создание jwt токена
func (s *SessionManager) CreateTokenByAccessSession(session AccessSession) (Token, error) {
	expiresAt := time.Now().Add(s.cfg.AccessTokenDuration)
	accessClaims := &AccessSessionClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  s.cfg.Audience,
			Issuer:    s.cfg.Issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
		AccessSession: session,
	}
	accessToken, err := s.jwtHelper.CreateToken(accessClaims)
	if err != nil {
		return Token{}, serviceerr.MakeErr(err, "s.jwtHelper.CreateToken")
	}

	return Token{
		Token:     accessToken,
		ExpiresAt: expiresAt,
	}, nil
}

// CreateTokenByRefreshSession создание refresh jwt токена
func (s *SessionManager) CreateTokenByRefreshSession(refresh RefreshSession) (Token, error) {
	expiresAt := time.Now().Add(s.cfg.AccessTokenDuration)
	accessClaims := &RefreshSessionClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  s.cfg.Audience,
			Issuer:    s.cfg.Issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
		RefreshSession: refresh,
	}
	accessToken, err := s.jwtHelper.CreateToken(accessClaims)
	if err != nil {
		return Token{}, serviceerr.MakeErr(err, "s.jwtHelper.CreateToken")
	}

	return Token{
		Token:     accessToken,
		ExpiresAt: expiresAt,
	}, nil
}

// GetAccessSessionByToken получить данные авторизованного пользователя по jwt токену
func (s *SessionManager) GetAccessSessionByToken(token string) (AccessSession, error) {
	claims := new(AccessSessionClaims)
	err := s.jwtHelper.Parse(token, claims)
	if err != nil {
		return AccessSession{}, server.ErrUnauthenticated(ErrSessionNotFound)
	}
	return claims.AccessSession, nil
}

// GetRefreshSessionByToken получить данные авторизованного пользователя по jwt токену
func (s *SessionManager) GetRefreshSessionByToken(token string) (RefreshSession, error) {
	claims := new(RefreshSessionClaims)
	err := s.jwtHelper.Parse(token, claims)
	if err != nil {
		return RefreshSession{}, server.ErrUnauthenticated(ErrSessionNotFound)
	}
	return claims.RefreshSession, nil
}
