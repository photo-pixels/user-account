package token_handler

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/server"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/photo-pixels/user-account/internal/handler"
	"github.com/photo-pixels/user-account/internal/user_case/form"
	desc "github.com/photo-pixels/user-account/pkg/gen/user_account"
)

// GetTokens получение списка токенов
func (h *TokenHandler) GetTokens(ctx context.Context, request *desc.GetTokensRequest) (*desc.GetTokensResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("userID is invalid: %w", err))
	}

	res, err := h.token.GetTokens(ctx, form.GetTokens{UserID: userID})
	if err != nil {
		return nil, handler.HandleError(err, "token.GetTokens")
	}

	items, err := mapTokens(res)
	if err != nil {
		return nil, handler.HandleError(err, "mapApiTokens")
	}

	return &desc.GetTokensResponse{
		Items: items,
	}, nil
}

// CreateToken создание нового токена
func (h *TokenHandler) CreateToken(ctx context.Context, request *desc.CreateTokenRequest) (*desc.CreateTokenResponse, error) {
	formReq, err := mapCreateToken(request)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("mapCreateToken: %w", err))
	}

	token, err := h.token.CreateToken(ctx, formReq)
	if err != nil {
		return nil, handler.HandleError(err, "token.CreateToken")
	}
	return &desc.CreateTokenResponse{
		Token: token,
	}, nil
}

// DeleteToken удаление токена
func (h *TokenHandler) DeleteToken(ctx context.Context, request *desc.DeleteTokenRequest) (*emptypb.Empty, error) {
	tokenID, err := uuid.Parse(request.TokenId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("tokenID is invalid: %w", err))
	}

	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("userID is invalid: %w", err))
	}

	if err = h.token.DeleteToken(ctx, form.DeleteToken{
		TokenID: tokenID,
		UserID:  userID,
	}); err != nil {
		return nil, handler.HandleError(err, "token.DeleteToken")
	}

	return &emptypb.Empty{}, nil
}

func (h *TokenHandler) GetToken(ctx context.Context, request *desc.GetTokenRequest) (*desc.GetTokenResponse, error) {
	res, err := h.token.GetToken(ctx, form.GetToken{Token: request.Token})
	if err != nil {
		return nil, handler.HandleError(err, "token.GetToken")
	}

	token := mapToken(res)
	if err != nil {
		return nil, handler.HandleError(err, "mapApiToken")
	}

	return &desc.GetTokenResponse{
		Token: token,
	}, nil
}
