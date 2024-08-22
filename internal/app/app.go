package app

import (
	"context"
	"fmt"

	"github.com/photo-pixels/platform/config"
	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/server"

	"github.com/photo-pixels/user-account/internal/service/codes"
	"github.com/photo-pixels/user-account/internal/service/jwt_helper"
	"github.com/photo-pixels/user-account/internal/service/password"
	"github.com/photo-pixels/user-account/internal/service/session_manager"
	"github.com/photo-pixels/user-account/internal/storage"
	"github.com/photo-pixels/user-account/internal/storage/pgrepo"
	"github.com/photo-pixels/user-account/internal/user_case/auth"
	"github.com/photo-pixels/user-account/internal/user_case/permission"
	"github.com/photo-pixels/user-account/internal/user_case/token"
	"github.com/photo-pixels/user-account/internal/user_case/user"
)

// App приложение
type App struct {
	cfgProvider config.Provider
	logger      log.Logger
	// cfg
	serverCfg server.Config
	// adapter
	storageAdapter *storage.Adapter
	// service
	jwtHelper             *jwt_helper.JwtHelper
	passwordService       *password.Service
	sessionManagerService *session_manager.SessionManager
	confirmCodeService    *codes.Service
	// user case
	authUserCase       *auth.Service
	permissionUserCase *permission.Service
	userUserCase       *user.Service
	tokenUserCase      *token.Service
}

// NewApp новое приложение
func NewApp(cfgProvider config.Provider) *App {
	return &App{cfgProvider: cfgProvider}
}

// Create создание сервисов
func (a *App) Create(ctx context.Context) error {
	var err error
	a.logger = log.NewLogger()

	pgCfg, err := a.getPgConnConfig()
	if err != nil {
		return fmt.Errorf("getPgConnConfig: %w", err)
	}
	pool, err := pgrepo.NewPgConn(ctx, pgCfg)
	if err != nil {
		return fmt.Errorf("newPgConn: %w", err)
	}

	a.storageAdapter = storage.NewStorageAdapter(
		a.logger,
		pool,
	)

	a.serverCfg, err = a.getServerConfig()
	if err != nil {
		return fmt.Errorf("getServerConfig: %w", err)
	}

	jwtHelperCfg, err := a.getJwtHelperConfig()
	if err != nil {
		return fmt.Errorf("getJwtHelperConfig: %w", err)
	}
	a.jwtHelper, err = jwt_helper.NewHelper(jwtHelperCfg)
	if err != nil {
		return fmt.Errorf("jwt_helper.NewHelper: %w", err)
	}

	sessionManagerCfg, err := a.getSessionManagerConfig()
	if err != nil {
		return fmt.Errorf("getSessionManagerConfig: %w", err)
	}
	a.sessionManagerService = session_manager.NewSessionManager(
		a.logger,
		sessionManagerCfg,
		a.jwtHelper,
	)

	a.confirmCodeService = codes.NewService(
		a.logger,
		a.storageAdapter,
	)

	authCfg, err := a.getAuthConfig()
	if err != nil {
		return fmt.Errorf("getAuthConfig: %w", err)
	}
	a.authUserCase = auth.NewService(
		a.logger,
		a.storageAdapter,
		authCfg,
		a.confirmCodeService,
		a.passwordService,
		a.sessionManagerService,
	)

	a.permissionUserCase = permission.NewService(
		a.logger,
		a.storageAdapter,
	)

	a.userUserCase = user.NewService(
		a.logger,
		a.storageAdapter,
	)

	a.tokenUserCase = token.NewService(
		a.logger,
		a.storageAdapter,
	)

	return nil
}

// GetLogger получить логер
func (a *App) GetLogger() log.Logger {
	return a.logger
}

// GetServerConfig конфигуратор сервера
func (a *App) GetServerConfig() server.Config {
	return a.serverCfg
}

// AuthUserCase юзеркейс авторизации
func (a *App) AuthUserCase() *auth.Service {
	return a.authUserCase
}

// PermissionUserCase юзеркейс пермисий и ролей
func (a *App) PermissionUserCase() *permission.Service {
	return a.permissionUserCase
}

// UserUserCase юзеркейс пользователей
func (a *App) UserUserCase() *user.Service {
	return a.userUserCase
}

// TokenUserCase юзеркейс токенов
func (a *App) TokenUserCase() *token.Service {
	return a.tokenUserCase
}
