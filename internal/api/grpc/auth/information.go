package auth

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/pkg/grpc/auth"
)

func (s *Server) Healthz(context.Context, *auth.HealthzRequest) (*auth.HealthzResponse, error) {
	return &auth.HealthzResponse{}, nil
}
