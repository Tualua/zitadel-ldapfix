package change

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
	change_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/change"
	"github.com/Tualua/zitadel-ldapfix/pkg/grpc/message"
)

func EventsToChangesPb(changes []*query.Event, assetAPIPrefix string) []*change_pb.Change {
	c := make([]*change_pb.Change, len(changes))
	for i, change := range changes {
		c[i] = EventToChangePb(change, assetAPIPrefix)
	}
	return c
}

func EventToChangePb(change *query.Event, assetAPIPrefix string) *change_pb.Change {
	return &change_pb.Change{
		ChangeDate:               timestamppb.New(change.CreationDate),
		EventType:                message.NewLocalizedEventType(change.Type),
		Sequence:                 change.Sequence,
		EditorId:                 change.Editor.ID,
		EditorDisplayName:        change.Editor.DisplayName,
		EditorPreferredLoginName: change.Editor.PreferedLoginName,
		EditorAvatarUrl:          domain.AvatarURL(assetAPIPrefix, change.Aggregate.ResourceOwner, change.Editor.AvatarKey),
		ResourceOwnerId:          change.Aggregate.ResourceOwner,
	}
}
