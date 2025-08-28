package view

import (
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/repository/project"
	"github.com/Tualua/zitadel-ldapfix/internal/zerrors"
)

func ProjectByIDQuery(id, instanceID string, latestSequence uint64) (*eventstore.SearchQueryBuilder, error) {
	if id == "" {
		return nil, zerrors.ThrowPreconditionFailed(nil, "EVENT-dke74", "Errors.Project.ProjectIDMissing")
	}
	return eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		InstanceID(instanceID).
		AwaitOpenTransactions().
		SequenceGreater(latestSequence).
		AddQuery().
		AggregateTypes(project.AggregateType).
		AggregateIDs(id).
		EventTypes(
			project.ProjectAddedType,
			project.ProjectChangedType,
			project.ProjectDeactivatedType,
			project.ProjectReactivatedType,
			project.ProjectRemovedType,
			project.OIDCConfigAddedType,
			project.ApplicationRemovedType,
		).
		Builder(), nil
}
