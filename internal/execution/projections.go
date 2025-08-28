package execution

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore/handler/v2"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
	"github.com/Tualua/zitadel-ldapfix/internal/query/projection"
	"github.com/Tualua/zitadel-ldapfix/internal/queue"
)

var (
	projections []*handler.Handler
)

func Register(
	ctx context.Context,
	executionsCustomConfig projection.CustomConfig,
	workerConfig WorkerConfig,
	queries *query.Queries,
	eventTypes []string,
	queue *queue.Queue,
) {
	queue.ShouldStart()
	projections = []*handler.Handler{
		NewEventHandler(ctx, projection.ApplyCustomConfig(executionsCustomConfig), eventTypes, eventstore.AggregateTypeFromEventType, queries, queue),
	}
	queue.AddWorkers(NewWorker(workerConfig))
}

func Start(ctx context.Context) {
	for _, projection := range projections {
		projection.Start(ctx)
	}
}
