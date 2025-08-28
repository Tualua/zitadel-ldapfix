package setup

import (
	"context"
	_ "embed"

	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/serviceping"
	"github.com/Tualua/zitadel-ldapfix/internal/v2/system"
)

type GenerateSystemID struct {
	eventstore *eventstore.Eventstore
}

func (mig *GenerateSystemID) Execute(ctx context.Context, _ eventstore.Event) error {
	id, err := serviceping.GenerateSystemID()
	if err != nil {
		return err
	}
	_, err = mig.eventstore.Push(ctx, system.NewIDGeneratedEvent(ctx, id))
	return err
}

func (mig *GenerateSystemID) String() string {
	return "60_generate_system_id"
}
