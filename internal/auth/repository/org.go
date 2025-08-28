package repository

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	iam_model "github.com/Tualua/zitadel-ldapfix/internal/iam/model"
)

type OrgRepository interface {
	GetMyPasswordComplexityPolicy(ctx context.Context) (*iam_model.PasswordComplexityPolicyView, error)
	GetLoginText(ctx context.Context, orgID string) ([]*domain.CustomText, error)
}
