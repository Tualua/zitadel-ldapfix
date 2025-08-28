package model

import (
	"github.com/Tualua/zitadel-ldapfix/internal/iam/model"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

func PasswordComplexityViewToModel(policy *query.PasswordComplexityPolicy) *model.PasswordComplexityPolicyView {
	return &model.PasswordComplexityPolicyView{
		AggregateID:  policy.ID,
		Sequence:     policy.Sequence,
		CreationDate: policy.CreationDate,
		ChangeDate:   policy.ChangeDate,
		MinLength:    policy.MinLength,
		HasLowercase: policy.HasLowercase,
		HasUppercase: policy.HasUppercase,
		HasSymbol:    policy.HasSymbol,
		HasNumber:    policy.HasNumber,
		Default:      policy.IsDefault,
	}
}
