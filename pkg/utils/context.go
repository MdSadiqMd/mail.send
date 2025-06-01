package utils

import (
	"context"

	"github.com/MdSadiqMd/mail.send/internal/models"
	logger "github.com/MdSadiqMd/mail.send/pkg/log"
)

type contextKey string

const userContextKey contextKey = "user"

var logContext = logger.New("context")

func SetUserInContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		logContext.Error("user not found in context")
		return nil, false
	}
	return user, false
}

func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	user, ok := GetUserFromContext(ctx)
	if !ok || user == nil {
		logContext.Error("user not found in context")
		return 0, false
	}
	return user.Id, true
}

func GetUserRoleFromContext(ctx context.Context) (string, bool) {
	user, ok := GetUserFromContext(ctx)
	if !ok || user == nil {
		logContext.Error("user not found in context")
		return "", false
	}
	return user.Role, true
}
