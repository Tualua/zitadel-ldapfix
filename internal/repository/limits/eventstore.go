package limits

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

func init() {
	eventstore.RegisterFilterEventMapper(AggregateType, SetEventType, SetEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, ResetEventType, ResetEventMapper)
}
