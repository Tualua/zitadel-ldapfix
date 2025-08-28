package channels

//go:generate mockgen -package mock -destination ./mock/channel.mock.go github.com/Tualua/zitadel-ldapfix/internal/notification/channels NotificationChannel
//go:generate mockgen -package mock -destination ./mock/message.mock.go github.com/Tualua/zitadel-ldapfix/internal/notification/channels Message
