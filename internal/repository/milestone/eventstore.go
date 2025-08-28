package milestone

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

var (
	ReachedEventMapper = eventstore.GenericEventMapper[ReachedEvent]
	PushedEventMapper  = eventstore.GenericEventMapper[PushedEvent]
)

func init() {
	eventstore.RegisterFilterEventMapper(AggregateType, ReachedEventType, ReachedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PushedEventType, PushedEventMapper)
}
