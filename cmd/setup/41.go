package setup

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/query/projection"
	"github.com/Tualua/zitadel-ldapfix/internal/repository/instance"
)

type FillFieldsForInstanceDomains struct {
	eventstore *eventstore.Eventstore
}

func (mig *FillFieldsForInstanceDomains) Execute(ctx context.Context, _ eventstore.Event) error {
	instances, err := mig.eventstore.InstanceIDs(
		ctx,
		eventstore.NewSearchQueryBuilder(eventstore.ColumnsInstanceIDs).
			OrderDesc().
			AddQuery().
			AggregateTypes("instance").
			EventTypes(instance.InstanceAddedEventType).
			Builder(),
	)
	if err != nil {
		return err
	}
	for _, instance := range instances {
		ctx := authz.WithInstanceID(ctx, instance)
		if err := projection.InstanceDomainFields.Trigger(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (mig *FillFieldsForInstanceDomains) String() string {
	return "repeatable_fill_fields_for_instance_domains"
}

func (f *FillFieldsForInstanceDomains) Check(lastRun map[string]interface{}) bool {
	return true
}
