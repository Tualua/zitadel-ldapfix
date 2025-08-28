package model

import es_models "github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"

type ProjectGrantMember struct {
	es_models.ObjectRoot
	GrantID string
	UserID  string
	Roles   []string
}
