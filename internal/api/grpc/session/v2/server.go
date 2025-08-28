package session

import (
	"net/http"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/api/grpc/server"
	"github.com/Tualua/zitadel-ldapfix/internal/command"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
	"github.com/Tualua/zitadel-ldapfix/pkg/grpc/session/v2"
	"github.com/Tualua/zitadel-ldapfix/pkg/grpc/session/v2/sessionconnect"
)

var _ sessionconnect.SessionServiceHandler = (*Server)(nil)

type Server struct {
	command *command.Commands
	query   *query.Queries

	checkPermission domain.PermissionCheck
}

type Config struct{}

func CreateServer(
	command *command.Commands,
	query *query.Queries,
	checkPermission domain.PermissionCheck,
) *Server {
	return &Server{
		command:         command,
		query:           query,
		checkPermission: checkPermission,
	}
}

func (s *Server) RegisterConnectServer(interceptors ...connect.Interceptor) (string, http.Handler) {
	return sessionconnect.NewSessionServiceHandler(s, connect.WithInterceptors(interceptors...))
}

func (s *Server) FileDescriptor() protoreflect.FileDescriptor {
	return session.File_zitadel_session_v2_session_service_proto
}

func (s *Server) AppName() string {
	return session.SessionService_ServiceDesc.ServiceName
}

func (s *Server) MethodPrefix() string {
	return session.SessionService_ServiceDesc.ServiceName
}

func (s *Server) AuthMethods() authz.MethodMapping {
	return session.SessionService_AuthMethods
}

func (s *Server) RegisterGateway() server.RegisterGatewayFunc {
	return session.RegisterSessionServiceHandler
}
