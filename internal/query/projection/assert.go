package projection

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/zerrors"
)

func assertEvent[T eventstore.Event](event eventstore.Event) (T, error) {
	e, ok := event.(T)
	if !ok {
		return e, zerrors.ThrowInvalidArgumentf(nil, "HANDL-1m9fS", "reduce.wrong.event.type %T", event)
	}
	return e, nil
}
