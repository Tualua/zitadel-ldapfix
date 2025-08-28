package start

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/zitadel/logging"

	"github.com/Tualua/zitadel-ldapfix/cmd/encryption"
	"github.com/Tualua/zitadel-ldapfix/cmd/hooks"
	"github.com/Tualua/zitadel-ldapfix/internal/actions"
	admin_es "github.com/Tualua/zitadel-ldapfix/internal/admin/repository/eventsourcing"
	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/api/http/middleware"
	"github.com/Tualua/zitadel-ldapfix/internal/api/oidc"
	"github.com/Tualua/zitadel-ldapfix/internal/api/saml"
	scim_config "github.com/Tualua/zitadel-ldapfix/internal/api/scim/config"
	"github.com/Tualua/zitadel-ldapfix/internal/api/ui/console"
	"github.com/Tualua/zitadel-ldapfix/internal/api/ui/login"
	auth_es "github.com/Tualua/zitadel-ldapfix/internal/auth/repository/eventsourcing"
	"github.com/Tualua/zitadel-ldapfix/internal/cache/connector"
	"github.com/Tualua/zitadel-ldapfix/internal/command"
	"github.com/Tualua/zitadel-ldapfix/internal/config/hook"
	"github.com/Tualua/zitadel-ldapfix/internal/config/network"
	"github.com/Tualua/zitadel-ldapfix/internal/config/systemdefaults"
	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/execution"
	"github.com/Tualua/zitadel-ldapfix/internal/id"
	"github.com/Tualua/zitadel-ldapfix/internal/logstore"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/handlers"
	"github.com/Tualua/zitadel-ldapfix/internal/query/projection"
	"github.com/Tualua/zitadel-ldapfix/internal/serviceping"
	static_config "github.com/Tualua/zitadel-ldapfix/internal/static/config"
	metrics "github.com/Tualua/zitadel-ldapfix/internal/telemetry/metrics/config"
	profiler "github.com/Tualua/zitadel-ldapfix/internal/telemetry/profiler/config"
	tracing "github.com/Tualua/zitadel-ldapfix/internal/telemetry/tracing/config"
)

type Config struct {
	Log                 *logging.Config
	Port                uint16
	ExternalPort        uint16
	ExternalDomain      string
	ExternalSecure      bool
	TLS                 network.TLS
	InstanceHostHeaders []string
	PublicHostHeaders   []string
	HTTP2HostHeader     string
	HTTP1HostHeader     string
	WebAuthNName        string
	Database            database.Config
	Caches              *connector.CachesConfig
	Tracing             tracing.Config
	Metrics             metrics.Config
	Profiler            profiler.Config
	Projections         projection.Config
	Notifications       handlers.WorkerConfig
	Executions          execution.WorkerConfig
	Auth                auth_es.Config
	Admin               admin_es.Config
	UserAgentCookie     *middleware.UserAgentCookieConfig
	OIDC                oidc.Config
	SAML                saml.Config
	SCIM                scim_config.Config
	Login               login.Config
	Console             console.Config
	AssetStorage        static_config.AssetStorageConfig
	InternalAuthZ       authz.Config
	SystemAuthZ         authz.Config
	SystemDefaults      systemdefaults.SystemDefaults
	EncryptionKeys      *encryption.EncryptionKeyConfig
	DefaultInstance     command.InstanceSetup
	AuditLogRetention   time.Duration
	SystemAPIUsers      map[string]*authz.SystemAPIUser
	CustomerPortal      string
	Machine             *id.Config
	Actions             *actions.Config
	Eventstore          *eventstore.Config
	LogStore            *logstore.Configs
	Quotas              *QuotasConfig
	Telemetry           *handlers.TelemetryPusherConfig
	ServicePing         *serviceping.Config
}

type QuotasConfig struct {
	Access struct {
		logstore.EmitterConfig  `mapstructure:",squash"`
		middleware.AccessConfig `mapstructure:",squash"`
	}
	Execution *logstore.EmitterConfig
}

func MustNewConfig(v *viper.Viper) *Config {
	config := new(Config)

	err := v.Unmarshal(config,
		viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
			hooks.SliceTypeStringDecode[*domain.CustomMessageText],
			hooks.SliceTypeStringDecode[authz.RoleMapping],
			hooks.MapTypeStringDecode[string, *authz.SystemAPIUser],
			hooks.MapHTTPHeaderStringDecode,
			database.DecodeHook(false),
			actions.HTTPConfigDecodeHook,
			hook.EnumHookFunc(authz.MemberTypeString),
			hooks.MapTypeStringDecode[domain.Feature, any],
			hooks.SliceTypeStringDecode[*command.SetQuota],
			hook.Base64ToBytesHookFunc(),
			hook.TagToLanguageHookFunc(),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.TextUnmarshallerHookFunc(),
		)),
	)
	logging.OnError(err).Fatal("unable to read config")

	err = config.Log.SetLogger()
	logging.OnError(err).Fatal("unable to set logger")

	err = config.Tracing.NewTracer()
	logging.OnError(err).Fatal("unable to set tracer")

	err = config.Metrics.NewMeter()
	logging.OnError(err).Fatal("unable to set meter")

	err = config.Profiler.NewProfiler()
	logging.OnError(err).Fatal("unable to set profiler")

	id.Configure(config.Machine)
	if config.Actions != nil {
		actions.SetHTTPConfig(&config.Actions.HTTP)
	}

	// Copy the global role permissions mappings to the instance until we allow instance-level configuration over the API.
	config.DefaultInstance.RolePermissionMappings = config.InternalAuthZ.RolePermissionMappings

	return config
}
