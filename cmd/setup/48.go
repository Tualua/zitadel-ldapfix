package setup

import (
	"context"
	_ "embed"

	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

var (
	//go:embed 48.sql
	addSAMLAppLoginVersion string
)

type Apps7SAMLConfigsLoginVersion struct {
	dbClient *database.DB
}

func (mig *Apps7SAMLConfigsLoginVersion) Execute(ctx context.Context, _ eventstore.Event) error {
	_, err := mig.dbClient.ExecContext(ctx, addSAMLAppLoginVersion)
	return err
}

func (mig *Apps7SAMLConfigsLoginVersion) String() string {
	return "48_apps7_saml_configs_login_version"
}
