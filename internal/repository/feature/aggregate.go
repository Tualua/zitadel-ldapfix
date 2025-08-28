package feature

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

const (
	eventTypePrefix = eventstore.EventType("feature.")
	setSuffix       = ".set"
)

const (
	AggregateType    = "feature"
	AggregateVersion = "v1"
)
