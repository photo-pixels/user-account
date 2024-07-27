package user_handler

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/server"

	"github.com/photo-pixels/user-account/internal/handler"
	desc "github.com/photo-pixels/user-account/pkg/gen/api/user_account"
)

// User получение данных пользователя
func (h *UserHandler) User(ctx context.Context, request *desc.UserRequest) (*desc.UserResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("userID is invalid: %w", err))
	}

	user, err := h.user.GetUser(ctx, userID)
	if err != nil {
		return nil, handler.HandleError(err, "h.user.GetUser")
	}

	return mapUserResponse(user)
}