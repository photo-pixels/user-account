package token_handler

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/server"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/photo-pixels/user-account/internal/user_case/dto"
	"github.com/photo-pixels/user-account/internal/user_case/form"
	desc "github.com/photo-pixels/user-account/pkg/gen/user_account"
)

func mapCreateToken(request *desc.CreateTokenRequest) (form.CreateToken, error) {
	var timeDuration *time.Duration
	if request.TimeDuration != nil {
		tt, err := time.ParseDuration(request.GetTimeDuration())
		if err != nil {
			return form.CreateToken{}, fmt.Errorf("invalid time duration: %s", request.GetTimeDuration())
		}
		timeDuration = &tt
	}

	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return form.CreateToken{}, server.ErrInvalidArgument(fmt.Errorf("userID is invalid: %w", err))
	}

	return form.CreateToken{
		Title:        request.Title,
		TokenType:    request.TokenType,
		UserID:       userID,
		TimeDuration: timeDuration,
	}, nil
}

func mapToken(item dto.Token) *desc.Token {
	var expiredAt *timestamppb.Timestamp
	if item.ExpiredAt != nil {
		expiredAt = timestamppb.New(*item.ExpiredAt)
	}
	return &desc.Token{
		Title:     item.Title,
		TokenType: item.TokenType,
		UserId:    item.UserID.String(),
		ExpiredAt: expiredAt,
	}
}

func mapTokens(res []dto.Token) ([]*desc.Token, error) {
	var result = make([]*desc.Token, 0, len(res))

	for _, item := range res {
		result = append(result, mapToken(item))
	}

	return result, nil
}
