package domain

import (
	"context"
	"time"

	"github.com/Tualua/zitadel-ldapfix/internal/crypto"
	es_models "github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
	"github.com/Tualua/zitadel-ldapfix/internal/telemetry/tracing"
	"github.com/Tualua/zitadel-ldapfix/internal/zerrors"
)

type Password struct {
	es_models.ObjectRoot

	SecretString   string
	EncodedSecret  string
	ChangeRequired bool
}

func NewPassword(password string) *Password {
	return &Password{
		SecretString: password,
	}
}

type PasswordCode struct {
	es_models.ObjectRoot

	Code             *crypto.CryptoValue
	Expiry           time.Duration
	NotificationType NotificationType
}

func (p *Password) HashPasswordIfExisting(ctx context.Context, policy *PasswordComplexityPolicy, hasher *crypto.Hasher) error {
	if p.SecretString == "" {
		return nil
	}
	if policy == nil {
		return zerrors.ThrowPreconditionFailed(nil, "DOMAIN-s8ifS", "Errors.User.PasswordComplexityPolicy.NotFound")
	}
	if err := policy.Check(p.SecretString); err != nil {
		return err
	}
	_, spanHash := tracing.NewNamedSpan(ctx, "passwap.Hash")
	encoded, err := hasher.Hash(p.SecretString)
	spanHash.EndWithError(err)
	if err != nil {
		return err
	}
	p.EncodedSecret = encoded
	return nil
}

func NewPasswordCode(passwordGenerator crypto.Generator) (*PasswordCode, error) {
	passwordCodeCrypto, _, err := crypto.NewCode(passwordGenerator)
	if err != nil {
		return nil, err
	}
	return &PasswordCode{
		Code:   passwordCodeCrypto,
		Expiry: passwordGenerator.Expiry(),
	}, nil
}
