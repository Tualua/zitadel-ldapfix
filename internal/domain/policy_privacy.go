package domain

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore/v1/models"
)

type PrivacyPolicy struct {
	models.ObjectRoot

	State   PolicyState
	Default bool

	TOSLink        string
	PrivacyLink    string
	HelpLink       string
	SupportEmail   EmailAddress
	DocsLink       string
	CustomLink     string
	CustomLinkText string
}
