package eventsourcing

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/authz/repository"
	authz_es "github.com/Tualua/zitadel-ldapfix/internal/authz/repository/eventsourcing/eventstore"
	authz_view "github.com/Tualua/zitadel-ldapfix/internal/authz/repository/eventsourcing/view"
	"github.com/Tualua/zitadel-ldapfix/internal/crypto"
	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

type EsRepository struct {
	authz_es.UserMembershipRepo
	authz_es.TokenVerifierRepo
}

func Start(queries *query.Queries, es *eventstore.Eventstore, dbClient *database.DB, keyEncryptionAlgorithm crypto.EncryptionAlgorithm, externalSecure bool) (repository.Repository, error) {
	view, err := authz_view.StartView(dbClient, queries)
	if err != nil {
		return nil, err
	}

	return &EsRepository{
		authz_es.UserMembershipRepo{
			Queries: queries,
		},
		authz_es.TokenVerifierRepo{
			TokenVerificationKey: keyEncryptionAlgorithm,
			Eventstore:           es,
			View:                 view,
			Query:                queries,
			ExternalSecure:       externalSecure,
		},
	}, nil
}

func (repo *EsRepository) Health(ctx context.Context) error {
	if err := repo.TokenVerifierRepo.Health(); err != nil {
		return err
	}
	return nil
}
