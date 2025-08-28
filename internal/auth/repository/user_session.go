package repository

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/user/model"
)

type UserSessionRepository interface {
	GetMyUserSessions(ctx context.Context) ([]*model.UserSessionView, error)
}
