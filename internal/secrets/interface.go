package secrets

import (
	"context"

	"github.com/spark-plus/autoloader/internal/connections"
)

type SecretManager interface {
	UpdateSecret(ctx context.Context, connectionID string, data *connections.ConnectionData) error
	DeleteSecret(ctx context.Context, connectionID string) error
}
