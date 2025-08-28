package types

import (
	"context"

	http_utils "github.com/Tualua/zitadel-ldapfix/internal/api/http"
	"github.com/Tualua/zitadel-ldapfix/internal/api/ui/login"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

func (notify Notify) SendUserInitCode(ctx context.Context, user *query.NotifyUser, code, authRequestID string) error {
	url := login.InitUserLink(http_utils.DomainContext(ctx).Origin(), user.ID, user.PreferredLoginName, code, user.ResourceOwner, user.PasswordSet, authRequestID)
	args := make(map[string]interface{})
	args["Code"] = code
	return notify(url, args, domain.InitCodeMessageType, true)
}
