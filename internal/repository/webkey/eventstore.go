package webkey

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

func init() {
	eventstore.RegisterFilterEventMapper(AggregateType, AddedEventType, eventstore.GenericEventMapper[AddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, ActivatedEventType, eventstore.GenericEventMapper[ActivatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, DeactivatedEventType, eventstore.GenericEventMapper[DeactivatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, RemovedEventType, eventstore.GenericEventMapper[RemovedEvent])
}
