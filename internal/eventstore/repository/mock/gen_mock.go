package mock

//go:generate mockgen -package mock -destination ./repository.mock.go github.com/Tualua/zitadel-ldapfix/internal/eventstore Querier,Pusher
