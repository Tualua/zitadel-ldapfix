package setup

import (
	"context"
	_ "embed"

	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

var (
	//go:embed 17.sql
	addOffsetField string
)

type AddOffsetToCurrentStates struct {
	dbClient *database.DB
}

func (mig *AddOffsetToCurrentStates) Execute(ctx context.Context, _ eventstore.Event) error {
	_, err := mig.dbClient.ExecContext(ctx, addOffsetField)
	return err
}

func (mig *AddOffsetToCurrentStates) String() string {
	return "17_add_offset_col_to_current_states"
}
