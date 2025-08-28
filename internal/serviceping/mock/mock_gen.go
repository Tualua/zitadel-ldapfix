package mock

//go:generate mockgen -package mock -destination queue.mock.go github.com/Tualua/zitadel-ldapfix/internal/serviceping Queue
//go:generate mockgen -package mock -destination queries.mock.go github.com/Tualua/zitadel-ldapfix/internal/serviceping Queries
//go:generate mockgen -package mock -destination telemetry.mock.go github.com/Tualua/zitadel-ldapfix/pkg/grpc/analytics/v2beta TelemetryServiceClient
