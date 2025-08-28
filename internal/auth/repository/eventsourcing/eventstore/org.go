package eventstore

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	auth_view "github.com/Tualua/zitadel-ldapfix/internal/auth/repository/eventsourcing/view"
	"github.com/Tualua/zitadel-ldapfix/internal/config/systemdefaults"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	iam_model "github.com/Tualua/zitadel-ldapfix/internal/iam/model"
	iam_view_model "github.com/Tualua/zitadel-ldapfix/internal/iam/repository/view/model"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

type OrgRepository struct {
	SearchLimit uint64

	Eventstore     *eventstore.Eventstore
	View           *auth_view.View
	SystemDefaults systemdefaults.SystemDefaults
	Query          *query.Queries
}

func (repo *OrgRepository) GetMyPasswordComplexityPolicy(ctx context.Context) (*iam_model.PasswordComplexityPolicyView, error) {
	policy, err := repo.Query.PasswordComplexityPolicyByOrg(ctx, false, authz.GetCtxData(ctx).OrgID, false)
	if err != nil {
		return nil, err
	}
	return iam_view_model.PasswordComplexityViewToModel(policy), err
}

func (repo *OrgRepository) GetLoginText(ctx context.Context, orgID string) ([]*domain.CustomText, error) {
	loginTexts, err := repo.Query.CustomTextListByTemplate(ctx, authz.GetInstance(ctx).InstanceID(), domain.LoginCustomText, false)
	if err != nil {
		return nil, err
	}
	orgLoginTexts, err := repo.Query.CustomTextListByTemplate(ctx, orgID, domain.LoginCustomText, false)
	if err != nil {
		return nil, err
	}
	return append(query.CustomTextsToDomain(loginTexts), query.CustomTextsToDomain(orgLoginTexts)...), nil
}
