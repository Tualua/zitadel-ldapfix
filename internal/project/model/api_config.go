package model

import (
	"github.com/Tualua/zitadel-ldapfix/internal/crypto"
	es_models "github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
)

type APIConfig struct {
	es_models.ObjectRoot
	AppID              string
	ClientID           string
	ClientSecret       *crypto.CryptoValue
	ClientSecretString string
	AuthMethodType     APIAuthMethodType
}

type APIAuthMethodType int32

const (
	APIAuthMethodTypeBasic APIAuthMethodType = iota
	APIAuthMethodTypePrivateKeyJWT
)
