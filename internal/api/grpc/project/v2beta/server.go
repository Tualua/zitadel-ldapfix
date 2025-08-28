package project

import (
	"net/http"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/command"
	"github.com/Tualua/zitadel-ldapfix/internal/config/systemdefaults"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
	project "github.com/Tualua/zitadel-ldapfix/pkg/grpc/project/v2beta"
	"github.com/Tualua/zitadel-ldapfix/pkg/grpc/project/v2beta/projectconnect"
)

var _ projectconnect.ProjectServiceHandler = (*Server)(nil)

type Server struct {
	systemDefaults systemdefaults.SystemDefaults
	command        *command.Commands
	query          *query.Queries

	checkPermission domain.PermissionCheck
}

type Config struct{}

func CreateServer(
	systemDefaults systemdefaults.SystemDefaults,
	command *command.Commands,
	query *query.Queries,
	checkPermission domain.PermissionCheck,
) *Server {
	return &Server{
		systemDefaults:  systemDefaults,
		command:         command,
		query:           query,
		checkPermission: checkPermission,
	}
}

func (s *Server) RegisterConnectServer(interceptors ...connect.Interceptor) (string, http.Handler) {
	return projectconnect.NewProjectServiceHandler(s, connect.WithInterceptors(interceptors...))
}

func (s *Server) FileDescriptor() protoreflect.FileDescriptor {
	return project.File_zitadel_project_v2beta_project_service_proto
}

func (s *Server) AppName() string {
	return project.ProjectService_ServiceDesc.ServiceName
}

func (s *Server) MethodPrefix() string {
	return project.ProjectService_ServiceDesc.ServiceName
}

func (s *Server) AuthMethods() authz.MethodMapping {
	return project.ProjectService_AuthMethods
}
