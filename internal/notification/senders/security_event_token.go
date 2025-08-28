package senders

import (
	"context"

	"github.com/zitadel/logging"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels/fs"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels/instrumenting"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels/log"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels/set"
)

const setSpanName = "security_event_token.NotificationChannel"

func SecurityEventTokenChannels(
	ctx context.Context,
	setConfig set.Config,
	getFileSystemProvider func(ctx context.Context) (*fs.Config, error),
	getLogProvider func(ctx context.Context) (*log.Config, error),
	successMetricName,
	failureMetricName string,
) (*Chain, error) {
	if err := setConfig.Validate(); err != nil {
		return nil, err
	}
	channels := make([]channels.NotificationChannel, 0, 3)
	setChannel, err := set.InitChannel(ctx, setConfig)
	logging.WithFields(
		"instance", authz.GetInstance(ctx).InstanceID(),
		"callurl", setConfig.CallURL,
	).OnError(err).Debug("initializing SET channel failed")
	if err == nil {
		channels = append(
			channels,
			instrumenting.Wrap(
				ctx,
				setChannel,
				setSpanName,
				successMetricName,
				failureMetricName,
			),
		)
	}
	channels = append(channels, debugChannels(ctx, getFileSystemProvider, getLogProvider)...)
	return ChainChannels(channels...), nil
}
