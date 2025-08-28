package admin

import (
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	admin_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/admin"
)

func UpdatePrivacyPolicyToDomain(req *admin_pb.UpdatePrivacyPolicyRequest) *domain.PrivacyPolicy {
	return &domain.PrivacyPolicy{
		TOSLink:        req.TosLink,
		PrivacyLink:    req.PrivacyLink,
		HelpLink:       req.HelpLink,
		SupportEmail:   domain.EmailAddress(req.SupportEmail),
		DocsLink:       req.DocsLink,
		CustomLink:     req.CustomLink,
		CustomLinkText: req.CustomLinkText,
	}
}
