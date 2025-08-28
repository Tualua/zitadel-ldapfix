package user

import (
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/v2/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/zerrors"
)

type machineAddedPayload struct {
	Username        string               `json:"userName"`
	Name            string               `json:"name,omitempty"`
	Description     string               `json:"description,omitempty"`
	AccessTokenType domain.OIDCTokenType `json:"accessTokenType,omitempty"`
}

type MachineAddedEvent eventstore.Event[machineAddedPayload]

const MachineAddedType = machinePrefix + ".added"

var _ eventstore.TypeChecker = (*MachineAddedEvent)(nil)

// ActionType implements eventstore.Typer.
func (c *MachineAddedEvent) ActionType() string {
	return MachineAddedType
}

func MachineAddedEventFromStorage(event *eventstore.StorageEvent) (e *MachineAddedEvent, _ error) {
	if event.Type != e.ActionType() {
		return nil, zerrors.ThrowInvalidArgument(nil, "USER-WLLoW", "Errors.Invalid.Event.Type")
	}

	payload, err := eventstore.UnmarshalPayload[machineAddedPayload](event.Payload)
	if err != nil {
		return nil, err
	}

	return &MachineAddedEvent{
		StorageEvent: event,
		Payload:      payload,
	}, nil
}
