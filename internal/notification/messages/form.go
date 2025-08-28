package messages

import (
	"net/url"

	"github.com/zitadel/schema"

	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels"
)

var _ channels.Message = (*Form)(nil)

type Form struct {
	Serializable        any
	TriggeringEventType eventstore.EventType
}

func (msg *Form) GetContent() (string, error) {
	values := make(url.Values)
	err := schema.NewEncoder().Encode(msg.Serializable, values)
	return values.Encode(), err
}

func (msg *Form) GetTriggeringEventType() eventstore.EventType {
	return msg.TriggeringEventType
}
