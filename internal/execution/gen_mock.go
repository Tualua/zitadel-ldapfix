package execution

//go:generate mockgen -package mock -destination ./mock/queries.mock.go github.com/Tualua/zitadel-ldapfix/internal/execution Queries
//go:generate mockgen -package mock -destination ./mock/queue.mock.go github.com/Tualua/zitadel-ldapfix/internal/execution Queue
