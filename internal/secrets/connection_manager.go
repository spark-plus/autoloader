package secrets

import (
	"context"
	"fmt"

	"github.com/sparkster/autoloader/internal/connections"
)

type ConnectionService struct {
	secretManager  SecretManager
	connectionRepo connections.ConnectionRepository
}

func (s *ConnectionService) InsertConnection(ctx context.Context, connection *connections.Connection, data *connections.ConnectionData) error {
	tx, err := s.connectionRepo.BeginTx(ctx)
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

	err = s.secretManager.UpdateSecret(ctx, connection.ConnectionID, data)
	if err != nil {
		return err
	}

	err = s.connectionRepo.CreateWTx(ctx, connection, tx)
	if err == nil {
		return nil
	}
	// rollback on failure

	err = s.secretManager.DeleteSecret(ctx, connection.ConnectionID)
	if err != nil {
		return fmt.Errorf("unable to rollback connectionID failed to create reference : %s", connection.ConnectionID)
	}

	return nil
}

func (s *ConnectionService) UpdateConnection(ctx context.Context, connection *connections.Connection, data *connections.ConnectionData) error {
	tx, err := s.connectionRepo.BeginTx(ctx)
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

	err = s.secretManager.UpdateSecret(ctx, connection.ConnectionID, data)
	if err != nil {
		return err
	}

	err = s.connectionRepo.UpdateWTx(ctx, connection, tx)
	if err == nil {
		return nil
	}
	// rollback on failure

	err = s.secretManager.DeleteSecret(ctx, connection.ConnectionID)
	if err != nil {
		return fmt.Errorf("unable to rollback connectionID failed to update reference : %s", connection.ConnectionID)
	}

	return nil
}
