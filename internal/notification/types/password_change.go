package types

import (
	"context"

	http_utils "github.com/Tualua/zitadel-ldapfix/internal/api/http"
	"github.com/Tualua/zitadel-ldapfix/internal/api/ui/console"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

func (notify Notify) SendPasswordChange(ctx context.Context, user *query.NotifyUser) error {
	url := console.LoginHintLink(http_utils.DomainContext(ctx).Origin(), user.PreferredLoginName)
	args := make(map[string]interface{})
	return notify(url, args, domain.PasswordChangeMessageType, true)
}
