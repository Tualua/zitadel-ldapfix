package policy

import (
	"github.com/Tualua/zitadel-ldapfix/internal/api/grpc/object"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
	policy_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/policy"
)

func ModelPrivacyPolicyToPb(policy *query.PrivacyPolicy) *policy_pb.PrivacyPolicy {
	return &policy_pb.PrivacyPolicy{
		IsDefault:    policy.IsDefault,
		TosLink:      policy.TOSLink,
		PrivacyLink:  policy.PrivacyLink,
		HelpLink:     policy.HelpLink,
		SupportEmail: string(policy.SupportEmail),
		Details: object.ToViewDetailsPb(
			policy.Sequence,
			policy.CreationDate,
			policy.ChangeDate,
			policy.ResourceOwner,
		),
		DocsLink:       policy.DocsLink,
		CustomLink:     policy.CustomLink,
		CustomLinkText: policy.CustomLinkText,
	}
}
