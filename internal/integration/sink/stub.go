//go:build !integration

package sink

import "github.com/Tualua/zitadel-ldapfix/internal/command"

// StartServer and its returned close function are a no-op
// when the `integration` build tag is disabled.
func StartServer(cmd *command.Commands) (close func()) {
	return func() {}
}
