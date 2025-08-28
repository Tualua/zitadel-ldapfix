package user

import (
	"google.golang.org/grpc"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/api/grpc/server"
	"github.com/Tualua/zitadel-ldapfix/internal/command"
	user "github.com/Tualua/zitadel-ldapfix/pkg/grpc/resources/user/v3alpha"
)

var _ user.ZITADELUsersServer = (*Server)(nil)

type Server struct {
	user.UnimplementedZITADELUsersServer
	command *command.Commands
}

type Config struct{}

func CreateServer(
	command *command.Commands,
) *Server {
	return &Server{
		command: command,
	}
}

func (s *Server) RegisterServer(grpcServer *grpc.Server) {
	user.RegisterZITADELUsersServer(grpcServer, s)
}

func (s *Server) AppName() string {
	return user.ZITADELUsers_ServiceDesc.ServiceName
}

func (s *Server) MethodPrefix() string {
	return user.ZITADELUsers_ServiceDesc.ServiceName
}

func (s *Server) AuthMethods() authz.MethodMapping {
	return user.ZITADELUsers_AuthMethods
}

func (s *Server) RegisterGateway() server.RegisterGatewayFunc {
	return user.RegisterZITADELUsersHandler
}
