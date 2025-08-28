package admin

import (
	_ "embed"
	"errors"

	"github.com/spf13/cobra"

	"github.com/Tualua/zitadel-ldapfix/cmd/initialise"
	"github.com/Tualua/zitadel-ldapfix/cmd/key"
	"github.com/Tualua/zitadel-ldapfix/cmd/setup"
	"github.com/Tualua/zitadel-ldapfix/cmd/start"
)

func New() *cobra.Command {
	adminCMD := &cobra.Command{
		Use:        "admin",
		Short:      "The ZITADEL admin CLI lets you interact with your instance",
		Long:       `The ZITADEL admin CLI lets you interact with your instance`,
		Deprecated: "please use subcommands directly, e.g. `zitadel start`",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no additional command provided")
		},
	}

	adminCMD.AddCommand(
		initialise.New(),
		setup.New(),
		start.New(nil),
		start.NewStartFromInit(nil),
		key.New(),
	)

	return adminCMD
}
