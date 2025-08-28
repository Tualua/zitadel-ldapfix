package domain

import (
	es_models "github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
)

type ProjectGrantMember struct {
	es_models.ObjectRoot

	GrantID string
	UserID  string
	Roles   []string
}

func (i *ProjectGrantMember) IsValid() bool {
	return i.AggregateID != "" && i.GrantID != "" && i.UserID != "" && len(i.Roles) != 0
}
