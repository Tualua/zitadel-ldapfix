package instance

import "github.com/Tualua/zitadel-ldapfix/internal/repository/instance"

const (
	AggregateType   = string(instance.AggregateType)
	eventTypePrefix = AggregateType + "."
)
