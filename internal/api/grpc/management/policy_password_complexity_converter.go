package management

import (
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	mgmt_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/management"
)

func AddPasswordComplexityPolicyToDomain(req *mgmt_pb.AddCustomPasswordComplexityPolicyRequest) *domain.PasswordComplexityPolicy {
	return &domain.PasswordComplexityPolicy{
		MinLength:    req.MinLength,
		HasLowercase: req.HasLowercase,
		HasUppercase: req.HasUppercase,
		HasNumber:    req.HasNumber,
		HasSymbol:    req.HasSymbol,
	}
}

func UpdatePasswordComplexityPolicyToDomain(req *mgmt_pb.UpdateCustomPasswordComplexityPolicyRequest) *domain.PasswordComplexityPolicy {
	return &domain.PasswordComplexityPolicy{
		MinLength:    req.MinLength,
		HasLowercase: req.HasLowercase,
		HasUppercase: req.HasUppercase,
		HasNumber:    req.HasNumber,
		HasSymbol:    req.HasSymbol,
	}
}
