//go:build integration

package settings_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Tualua/zitadel-ldapfix/internal/integration"
	settings "github.com/Tualua/zitadel-ldapfix/pkg/grpc/settings/v2beta"
)

var (
	CTX, AdminCTX context.Context
	Instance      *integration.Instance
	Client        settings.SettingsServiceClient
)

func TestMain(m *testing.M) {
	os.Exit(func() int {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
		defer cancel()

		Instance = integration.NewInstance(ctx)

		CTX = ctx
		AdminCTX = Instance.WithAuthorization(ctx, integration.UserTypeIAMOwner)
		Client = Instance.Client.SettingsV2beta
		return m.Run()
	}())
}
