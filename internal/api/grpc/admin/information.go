package admin

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/pkg/grpc/admin"
)

func (s *Server) Healthz(context.Context, *admin.HealthzRequest) (*admin.HealthzResponse, error) {
	return &admin.HealthzResponse{}, nil
}
