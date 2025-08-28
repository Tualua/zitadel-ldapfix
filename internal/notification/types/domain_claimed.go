package types

import (
	"context"
	"strings"

	http_utils "github.com/Tualua/zitadel-ldapfix/internal/api/http"
	"github.com/Tualua/zitadel-ldapfix/internal/api/ui/login"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

func (notify Notify) SendDomainClaimed(ctx context.Context, user *query.NotifyUser, username string) error {
	url := login.LoginLink(http_utils.DomainContext(ctx).Origin(), user.ResourceOwner)
	index := strings.LastIndex(user.LastEmail, "@")
	args := make(map[string]interface{})
	args["TempUsername"] = username
	args["Domain"] = user.LastEmail[index+1:]
	return notify(url, args, domain.DomainClaimedMessageType, true)
}
