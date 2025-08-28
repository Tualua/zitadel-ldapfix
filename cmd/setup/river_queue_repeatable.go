package setup

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/queue"
)

type RiverMigrateRepeatable struct {
	client *database.DB
}

func (mig *RiverMigrateRepeatable) Execute(ctx context.Context, _ eventstore.Event) error {
	return queue.NewMigrator(mig.client).Execute(ctx)
}

func (mig *RiverMigrateRepeatable) String() string {
	return "repeatable_migrate_river"
}

func (f *RiverMigrateRepeatable) Check(lastRun map[string]interface{}) bool {
	return true
}
