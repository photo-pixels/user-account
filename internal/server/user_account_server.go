package server

import (
	"context"

	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/server"

	"github.com/photo-pixels/user-account/internal/handler/user_account/auth_handler"
	"github.com/photo-pixels/user-account/internal/handler/user_account/permission_handler"
	"github.com/photo-pixels/user-account/internal/handler/user_account/user_handler"
)

// UserAccountServer сервер
type UserAccountServer struct {
	*CustomServer
}

// NewUserAccountServer новый сервер
func NewUserAccountServer(
	logger log.Logger,
	cfg server.Config,
	authUserCase auth_handler.AuthUserCase,
	permission permission_handler.PermissionUserCase,
	userUserCase user_handler.UserUserCase,
) *UserAccountServer {
	return &UserAccountServer{
		CustomServer: NewCustomServer(
			logger,
			cfg,
			auth_handler.NewHandler(logger, authUserCase),
			permission_handler.NewHandler(logger, permission),
			user_handler.NewHandler(logger, userUserCase),
		),
	}
}

// Start старт сервера
func (s *UserAccountServer) Start(ctx context.Context) error {
	return s.CustomServer.Start(ctx, "api")
}
