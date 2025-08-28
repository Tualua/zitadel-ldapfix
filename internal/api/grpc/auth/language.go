package auth

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/i18n"
	auth_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/auth"
)

func (s *Server) GetSupportedLanguages(context.Context, *auth_pb.GetSupportedLanguagesRequest) (*auth_pb.GetSupportedLanguagesResponse, error) {
	return &auth_pb.GetSupportedLanguagesResponse{Languages: domain.LanguagesToStrings(i18n.SupportedLanguages())}, nil
}
