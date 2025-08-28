package admin

import (
	"google.golang.org/protobuf/types/known/durationpb"

	obj_grpc "github.com/Tualua/zitadel-ldapfix/internal/api/grpc/object"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
	admin_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/admin"
	settings_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/settings"
)

func OIDCSettingsToPb(config *query.OIDCSettings) *settings_pb.OIDCSettings {
	return &settings_pb.OIDCSettings{
		Details:                    obj_grpc.ToViewDetailsPb(config.Sequence, config.CreationDate, config.ChangeDate, config.AggregateID),
		AccessTokenLifetime:        durationpb.New(config.AccessTokenLifetime),
		IdTokenLifetime:            durationpb.New(config.IdTokenLifetime),
		RefreshTokenIdleExpiration: durationpb.New(config.RefreshTokenIdleExpiration),
		RefreshTokenExpiration:     durationpb.New(config.RefreshTokenExpiration),
	}
}

func AddOIDCConfigToConfig(req *admin_pb.AddOIDCSettingsRequest) *domain.OIDCSettings {
	return &domain.OIDCSettings{
		AccessTokenLifetime:        req.AccessTokenLifetime.AsDuration(),
		IdTokenLifetime:            req.IdTokenLifetime.AsDuration(),
		RefreshTokenIdleExpiration: req.RefreshTokenIdleExpiration.AsDuration(),
		RefreshTokenExpiration:     req.RefreshTokenExpiration.AsDuration(),
	}
}

func UpdateOIDCConfigToConfig(req *admin_pb.UpdateOIDCSettingsRequest) *domain.OIDCSettings {
	return &domain.OIDCSettings{
		AccessTokenLifetime:        req.AccessTokenLifetime.AsDuration(),
		IdTokenLifetime:            req.IdTokenLifetime.AsDuration(),
		RefreshTokenIdleExpiration: req.RefreshTokenIdleExpiration.AsDuration(),
		RefreshTokenExpiration:     req.RefreshTokenExpiration.AsDuration(),
	}
}
