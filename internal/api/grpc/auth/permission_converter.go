package auth

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/api/grpc/object"
	user_grpc "github.com/Tualua/zitadel-ldapfix/internal/api/grpc/user"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
	auth_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/auth"
)

func ListMyMembershipsRequestToModel(ctx context.Context, req *auth_pb.ListMyMembershipsRequest) (*query.MembershipSearchQuery, error) {
	offset, limit, asc := object.ListQueryToModel(req.Query)
	queries, err := user_grpc.MembershipQueriesToQuery(req.Queries)
	if err != nil {
		return nil, err
	}
	userQuery, err := query.NewMembershipUserIDQuery(authz.GetCtxData(ctx).UserID)
	if err != nil {
		return nil, err
	}
	queries = append(queries, userQuery)
	return &query.MembershipSearchQuery{
		SearchRequest: query.SearchRequest{
			Offset: offset,
			Limit:  limit,
			Asc:    asc,
			//SortingColumn: //TODO: sorting
		},
		Queries: queries,
	}, nil
}
