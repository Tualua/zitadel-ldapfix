package email

import (
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels/smtp"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels/webhook"
)

type Config struct {
	ProviderConfig *Provider
	SMTPConfig     *smtp.Config
	WebhookConfig  *webhook.Config
}

type Provider struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}
