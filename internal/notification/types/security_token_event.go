package types

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels/set"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/messages"
)

func handleSecurityTokenEvent(
	ctx context.Context,
	setConfig set.Config,
	channels ChannelChains,
	token any,
	triggeringEventType eventstore.EventType,
) error {
	message := &messages.Form{
		Serializable:        token,
		TriggeringEventType: triggeringEventType,
	}
	setChannels, err := channels.SecurityTokenEvent(ctx, setConfig)
	if err != nil {
		return err
	}
	return setChannels.HandleMessage(message)
}
