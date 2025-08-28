package admin

import (
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	admin_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/admin"
)

func UpdatePasswordAgePolicyToDomain(policy *admin_pb.UpdatePasswordAgePolicyRequest) *domain.PasswordAgePolicy {
	return &domain.PasswordAgePolicy{
		MaxAgeDays:     uint64(policy.MaxAgeDays),
		ExpireWarnDays: uint64(policy.ExpireWarnDays),
	}
}
