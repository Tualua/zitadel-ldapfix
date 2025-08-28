package setup

import (
	"context"
	_ "embed"

	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

var (
	//go:embed 28.sql
	addFieldTable string
)

type AddFieldTable struct {
	dbClient *database.DB
}

func (mig *AddFieldTable) Execute(ctx context.Context, _ eventstore.Event) error {
	_, err := mig.dbClient.ExecContext(ctx, addFieldTable)
	return err
}

func (mig *AddFieldTable) String() string {
	return "28_add_search_table"
}
