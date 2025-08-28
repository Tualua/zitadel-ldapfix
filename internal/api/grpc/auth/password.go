package auth

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/api/grpc/object"
	auth_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/auth"
)

func (s *Server) UpdateMyPassword(ctx context.Context, req *auth_pb.UpdateMyPasswordRequest) (*auth_pb.UpdateMyPasswordResponse, error) {
	ctxData := authz.GetCtxData(ctx)
	objectDetails, err := s.command.ChangePassword(ctx, ctxData.ResourceOwner, ctxData.UserID, req.OldPassword, req.NewPassword, "", false)
	if err != nil {
		return nil, err
	}
	return &auth_pb.UpdateMyPasswordResponse{
		Details: object.DomainToChangeDetailsPb(objectDetails),
	}, nil
}
