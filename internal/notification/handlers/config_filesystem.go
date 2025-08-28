package handlers

import (
	"context"

	"github.com/Tualua/zitadel-ldapfix/internal/api/authz"
	"github.com/Tualua/zitadel-ldapfix/internal/domain"
	"github.com/Tualua/zitadel-ldapfix/internal/notification/channels/fs"
)

// GetFileSystemProvider reads the iam filesystem provider config
func (n *NotificationQueries) GetFileSystemProvider(ctx context.Context) (*fs.Config, error) {
	config, err := n.NotificationProviderByIDAndType(ctx, authz.GetInstance(ctx).InstanceID(), domain.NotificationProviderTypeFile)
	if err != nil {
		return nil, err
	}
	return &fs.Config{
		Compact: config.Compact,
		Path:    n.fileSystemPath,
	}, nil
}
