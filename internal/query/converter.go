package query

import (
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
)

func readModelToObjectDetails(model *eventstore.ReadModel) *domain.ObjectDetails {
	return &domain.ObjectDetails{
		Sequence:      model.ProcessedSequence,
		ResourceOwner: model.ResourceOwner,
		EventDate:     model.ChangeDate,
	}
}
