package connect_middleware

import (
	"context"

	"connectrpc.com/connect"
	"github.com/zitadel/logging"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/i18n"
	_ "github.com/Tualua/zitadel-ldapfix/internal/statik"
	"github.com/Tualua/zitadel-ldapfix/internal/telemetry/tracing"
)

func TranslationHandler() connect.UnaryInterceptorFunc {

	return func(handler connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := handler(ctx, req)
			ctx, span := tracing.NewSpan(ctx)
			defer func() { span.EndWithError(err) }()

			if err != nil {
				translator, translatorError := getTranslator(ctx)
				if translatorError != nil {
					return resp, err
				}
				return resp, translateError(ctx, err, translator)
			}
			if loc, ok := resp.Any().(localizers); ok {
				translator, translatorError := getTranslator(ctx)
				if translatorError != nil {
					return resp, err
				}
				translateFields(ctx, loc, translator)
			}
			return resp, nil
		}
	}
}

func getTranslator(ctx context.Context) (*i18n.Translator, error) {
	translator, err := i18n.NewZitadelTranslator(authz.GetInstance(ctx).DefaultLanguage())
	if err != nil {
		logging.New().WithError(err).Error("could not load translator")
	}
	return translator, err
}
