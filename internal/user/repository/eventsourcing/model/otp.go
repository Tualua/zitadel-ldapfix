package model

import (
	"github.com/zitadel/logging"

	"github.com/Tualua/zitadel-ldapfix/internal/crypto"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	es_models "github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
	"github.com/Tualua/zitadel-ldapfix/internal/user/model"
	"github.com/Tualua/zitadel-ldapfix/internal/zerrors"
)

type OTP struct {
	es_models.ObjectRoot

	Secret *crypto.CryptoValue `json:"otpSecret,omitempty"`
	State  int32               `json:"-"`
}

type OTPVerified struct {
	UserAgentID string `json:"userAgentID,omitempty"`
}

func (u *Human) appendOTPAddedEvent(event eventstore.Event) error {
	u.OTP = &OTP{
		State: int32(model.MFAStateNotReady),
	}
	return u.OTP.setData(event)
}

func (u *Human) appendOTPVerifiedEvent() {
	u.OTP.State = int32(model.MFAStateReady)
}

func (u *Human) appendOTPRemovedEvent() {
	u.OTP = nil
}

func (o *OTP) setData(event eventstore.Event) error {
	o.ObjectRoot.AppendEvent(event)
	if err := event.Unmarshal(o); err != nil {
		logging.Log("EVEN-d9soe").WithError(err).Error("could not unmarshal event data")
		return zerrors.ThrowInternal(err, "MODEL-lo023", "could not unmarshal event")
	}
	return nil
}

func (o *OTPVerified) SetData(event eventstore.Event) error {
	if err := event.Unmarshal(o); err != nil {
		logging.Log("EVEN-BF421").WithError(err).Error("could not unmarshal event data")
		return zerrors.ThrowInternal(err, "MODEL-GB6hj", "could not unmarshal event")
	}
	return nil
}
