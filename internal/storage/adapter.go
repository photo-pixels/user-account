package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/log"

	"github.com/photo-pixels/user-account/internal/model"
)

// Adapter адаптер для работы с базой
type Adapter struct {
	logger log.Logger
}

// NewStorageAdapter новый адаптер
func NewStorageAdapter(
	logger log.Logger,
) *Adapter {
	return &Adapter{
		logger: logger.Named("storage_adapter"),
	}
}

func (a *Adapter) CreatePermission(ctx context.Context, permission model.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) AddPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) CreateRole(ctx context.Context, role model.Role) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) AddRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetPermissions(ctx context.Context, params model.GetPermissions) ([]model.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetRoles(ctx context.Context, params model.GetRoles) ([]model.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]model.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetLastActiveRefreshToken(ctx context.Context, refreshTokenID uuid.UUID) (model.RefreshToken, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetUser(ctx context.Context, userID uuid.UUID) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) SaveConfirmCode(ctx context.Context, confirmCode model.ConfirmCode) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetActiveConfirmCode(ctx context.Context, code string, confirmType model.ConfirmCodeType) (model.ConfirmCode, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) UpdateConfirmCode(ctx context.Context, personID uuid.UUID, confirmCodeType model.ConfirmCodeType, update model.UpdateConfirmCode) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) RunTransaction(ctx context.Context, txFunc func(ctxTx context.Context) error) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) EmailExists(ctx context.Context, email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) SaveUser(ctx context.Context, user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) SaveUserAuth(ctx context.Context, auth model.Auth) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetAuth(ctx context.Context, userID uuid.UUID) (model.Auth, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) UpdateUser(ctx context.Context, userID uuid.UUID, updateUser model.UpdateUser) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) UpdateUserAuth(ctx context.Context, userID uuid.UUID, updateAuth model.UpdateAuth) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetAuthByEmail(ctx context.Context, email string) (model.Auth, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) SaveRefreshToken(ctx context.Context, refreshToken model.RefreshToken) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) UpdateRefreshTokenStatus(ctx context.Context, refreshTokenID uuid.UUID, status model.RefreshTokenStatus) error {
	//TODO implement me
	panic("implement me")
}
