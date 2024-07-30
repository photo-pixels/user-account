package permission_handler

import (
	"context"
	"net/http"

	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/server"
	"google.golang.org/grpc"

	"github.com/photo-pixels/user-account/internal/user_case/dto"
	"github.com/photo-pixels/user-account/internal/user_case/form"
	desc "github.com/photo-pixels/user-account/pkg/gen/api/user_account"
)

// PermissionUserCase интерфейс юзер кейса работы с пермиссиями
type PermissionUserCase interface {
	CreateRole(ctx context.Context, form form.CreateRole) (dto.Role, error)
	CreatePermission(ctx context.Context, form form.CreatePermission) (dto.Permission, error)
	AddPermissionToRole(ctx context.Context, form form.AddPermissionToRole) error
	GetUserPermissions(ctx context.Context, form form.GetUserPermissions) ([]dto.Permission, error)
	AddRoleToUser(ctx context.Context, form form.AddRoleToUser) error
}

// PermissionHandler хендлер работы с пермиссиями
type PermissionHandler struct {
	desc.PermissionServiceServer
	logger     log.Logger
	permission PermissionUserCase
}

// NewHandler новый хендлер работы с пермиссиями
func NewHandler(logger log.Logger,
	permission PermissionUserCase,
) *PermissionHandler {
	return &PermissionHandler{
		logger:     logger.Named("permission_handler"),
		permission: permission,
	}
}

// RegistrationServerHandlers .
func (h *PermissionHandler) RegistrationServerHandlers(_ *http.ServeMux) {
}

// RegisterServiceHandlerFromEndpoint .
func (h *PermissionHandler) RegisterServiceHandlerFromEndpoint() server.HandlerFromEndpoint {
	return desc.RegisterPermissionServiceHandlerFromEndpoint
}

// RegisterServiceServer .
func (h *PermissionHandler) RegisterServiceServer(server *grpc.Server) {
	desc.RegisterPermissionServiceServer(server, h)
}
