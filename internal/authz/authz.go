package authz

import (
	"github.com/Tualua/zitadel-ldapfix/internal/authz/repository"
	"github.com/Tualua/zitadel-ldapfix/internal/authz/repository/eventsourcing"
	"github.com/Tualua/zitadel-ldapfix/internal/crypto"
	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

func Start(queries *query.Queries, es *eventstore.Eventstore, dbClient *database.DB, keyEncryptionAlgorithm crypto.EncryptionAlgorithm, externalSecure bool) (repository.Repository, error) {
	return eventsourcing.Start(queries, es, dbClient, keyEncryptionAlgorithm, externalSecure)
}
