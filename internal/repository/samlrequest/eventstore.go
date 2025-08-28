package samlrequest

import "github.com/Tualua/zitadel-ldapfix/internal/eventstore"

func init() {
	eventstore.RegisterFilterEventMapper(AggregateType, AddedType, eventstore.GenericEventMapper[AddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SessionLinkedType, eventstore.GenericEventMapper[SessionLinkedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, FailedType, eventstore.GenericEventMapper[FailedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, SucceededType, eventstore.GenericEventMapper[SucceededEvent])
}
