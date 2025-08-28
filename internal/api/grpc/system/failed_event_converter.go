package system

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Tualua/zitadel-ldapfix/internal/query"
	system_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/system"
)

func FailedEventsToPb(database string, failedEvents *query.FailedEvents) []*system_pb.FailedEvent {
	events := make([]*system_pb.FailedEvent, len(failedEvents.FailedEvents))
	for i, failedEvent := range failedEvents.FailedEvents {
		events[i] = FailedEventToPb(database, failedEvent)
	}
	return events
}

func FailedEventToPb(database string, failedEvent *query.FailedEvent) *system_pb.FailedEvent {
	var lastFailed *timestamppb.Timestamp
	if !failedEvent.LastFailed.IsZero() {
		lastFailed = timestamppb.New(failedEvent.LastFailed)
	}
	return &system_pb.FailedEvent{
		Database:       database,
		ViewName:       failedEvent.ProjectionName,
		FailedSequence: failedEvent.FailedSequence,
		FailureCount:   failedEvent.FailureCount,
		ErrorMessage:   failedEvent.Error,
		LastFailed:     lastFailed,
	}
}
