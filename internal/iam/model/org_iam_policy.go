package model

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
)

type DomainPolicy struct {
	models.ObjectRoot

	State                 PolicyState
	UserLoginMustBeDomain bool
	Default               bool
}
