package domain

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
)

type PasswordAgePolicy struct {
	models.ObjectRoot

	MaxAgeDays     uint64
	ExpireWarnDays uint64
}
