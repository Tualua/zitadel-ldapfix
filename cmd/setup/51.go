package setup

import (
	"context"
	_ "embed"

	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

var (
	//go:embed 51.sql
	addRootCA string
)

type IDPTemplate6RootCA struct {
	dbClient *database.DB
}

func (mig *IDPTemplate6RootCA) Execute(ctx context.Context, _ eventstore.Event) error {
	_, err := mig.dbClient.ExecContext(ctx, addRootCA)
	return err
}

func (mig *IDPTemplate6RootCA) String() string {
	return "51_idp_templates6_add_root_ca"
}
