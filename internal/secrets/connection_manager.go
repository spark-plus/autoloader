package secrets

import (
	"context"

	"github.com/sparkster/autoloader/internal/connections"
)

type ConnectionService struct {
	secretManager  SecretManager
	connectionRepo connections.ConnectionRepository
}

func (s *ConnectionService) CreateConnection(ctx context.Context, connection *connections.Connection, data *connections.ConnectionData) error {
	tx, err := s.connectionRepo.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	err = s.secretManagerRepo.PutConnectionData(ctx, connection.ConnectionID, data)
	if err != nil {
		return err
	}

	err = s.connectionRepo.Create(ctx, connection)
	if err != nil {
		return err
	}

	return nil
}
