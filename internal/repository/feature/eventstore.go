package feature

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

func init() {
	eventstore.RegisterFilterEventMapper(AggregateType, DefaultLoginInstanceEventType, eventstore.GenericEventMapper[SetEvent[Boolean]])
}
