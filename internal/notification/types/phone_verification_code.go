package types

import (
	"context"

	http_util "github.com/Tualua/zitadel-ldapfix/internal/api/http"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
)

func (notify Notify) SendPhoneVerificationCode(ctx context.Context, code string) error {
	args := make(map[string]interface{})
	args["Code"] = code
	args["Domain"] = http_util.DomainContext(ctx).RequestedDomain()
	return notify("", args, domain.VerifyPhoneMessageType, true)
}
