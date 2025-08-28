package keypair

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

const (
	AggregateType    = "key_pair"
	AggregateVersion = "v1"
)

type Aggregate struct {
	eventstore.Aggregate
}
