package model

import (
	es_models "github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
)

type SAMLConfig struct {
	es_models.ObjectRoot
	AppID       string
	Metadata    []byte
	MetadataURL string
}
