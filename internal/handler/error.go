package handler

import (
	"errors"
	"fmt"

	"github.com/photo-pixels/platform/server"
	"github.com/photo-pixels/platform/serviceerr"

	desc "github.com/photo-pixels/user-account/pkg/gen/user_account"
)

// HandleError обработчик ошибок
func HandleError(err error, description any) error {
	newErr := fmt.Errorf("%s: %w", description, err)

	info := desc.ErrorInfo{
		Description: "Unhandled error",
	}

	var serviceErr *serviceerr.ErrorService
	if errors.As(newErr, &serviceErr) {
		info = desc.ErrorInfo{
			Description: serviceErr.ErrInfo.Description,
			// FieldViolations: mapFieldViolation(serviceErr.ErrInfo.FieldViolations),
		}
		switch serviceErr.Type {
		case serviceerr.InvalidInputDataErrorType:
			return server.ErrInvalidArgument(newErr, &info)
		case serviceerr.RuntimeErrorType:
			return server.ErrInternal(newErr, &info)
		case serviceerr.NotFoundErrorType:
			return server.ErrNotFound(newErr, &info)
		case serviceerr.ConflictErrorType:
			return server.ErrAlreadyExists(newErr, &info)
		case serviceerr.FailPreconditionErrorType:
			return server.ErrFailedPrecondition(newErr, &info)
		case serviceerr.PermissionDeniedType:
			return server.ErrPermissionDenied(newErr, &info)
		}
	}

	return server.ErrInternal(newErr, &info)
}
