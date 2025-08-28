package types

import (
	"context"
	"strings"

	http_utils "github.com/Tualua/zitadel-ldapfix/internal/api/http"
	"github.com/Tualua/zitadel-ldapfix/internal/api/ui/login"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

func (notify Notify) SendPasswordlessRegistrationLink(ctx context.Context, user *query.NotifyUser, code, codeID, urlTmpl string) error {
	var url string
	if urlTmpl == "" {
		url = domain.PasswordlessInitCodeLink(http_utils.DomainContext(ctx).Origin()+login.HandlerPrefix+login.EndpointPasswordlessRegistration, user.ID, user.ResourceOwner, codeID, code)
	} else {
		var buf strings.Builder
		if err := domain.RenderPasskeyURLTemplate(&buf, urlTmpl, user.ID, user.ResourceOwner, codeID, code); err != nil {
			return err
		}
		url = buf.String()
	}
	return notify(url, nil, domain.PasswordlessRegistrationMessageType, true)
}
