package domain

import (
	"time"

	"github.com/Tualua/zitadel-ldapfix/internal/crypto"
	es_models "github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
)

type KeyPair struct {
	es_models.ObjectRoot

	Usage       crypto.KeyUsage
	Algorithm   string
	PrivateKey  *Key
	PublicKey   *Key
	Certificate *Key
}

type Key struct {
	Key    *crypto.CryptoValue
	Expiry time.Time
}

func (k *KeyPair) IsValid() bool {
	return k.Algorithm != "" &&
		k.PrivateKey != nil && k.PrivateKey.IsValid() &&
		k.PublicKey != nil && k.PublicKey.IsValid() &&
		k.Certificate != nil && k.Certificate.IsValid()
}

func (k *Key) IsValid() bool {
	return k.Key != nil
}
