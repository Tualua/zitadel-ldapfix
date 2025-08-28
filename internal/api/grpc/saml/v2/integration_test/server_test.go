//go:build integration

package saml_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Tualua/zitadel-ldapfix/internal/integration"
	saml_pb "github.com/Tualua/zitadel-ldapfix/pkg/grpc/saml/v2"
)

var (
	CTX      context.Context
	IAMCTX   context.Context
	LoginCTX context.Context
	Instance *integration.Instance
	Client   saml_pb.SAMLServiceClient
)

func TestMain(m *testing.M) {
	os.Exit(func() int {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
		defer cancel()

		Instance = integration.NewInstance(ctx)
		Client = Instance.Client.SAMLv2

		IAMCTX = Instance.WithAuthorization(ctx, integration.UserTypeIAMOwner)
		CTX = Instance.WithAuthorization(ctx, integration.UserTypeOrgOwner)
		LoginCTX = Instance.WithAuthorization(ctx, integration.UserTypeLogin)
		return m.Run()
	}())
}
