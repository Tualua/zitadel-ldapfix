package setup

import (
	"context"
	_ "embed"

	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

var (
	//go:embed 19.sql
	addCurrentSequencesIndex string
)

type AddCurrentSequencesIndex struct {
	dbClient *database.DB
}

func (mig *AddCurrentSequencesIndex) Execute(ctx context.Context, _ eventstore.Event) error {
	_, err := mig.dbClient.ExecContext(ctx, addCurrentSequencesIndex)
	return err
}

func (mig *AddCurrentSequencesIndex) String() string {
	return "19_add_current_sequences_index"
}
