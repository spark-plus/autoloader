package secrets

import (
	"context"
	"fmt"

	"github.com/spark-plus/autoloader/internal/connections"
	"xorm.io/xorm"
)

type ConnectionService struct {
	secretManager SecretManager
	repo          connections.ConnectionRepo
}

func (s *ConnectionService) InsertConnection(ctx context.Context, connection *connections.Connection, data *connections.ConnectionData) error {
	err := s.repo.T().DoTransaction(ctx, func(session *xorm.Session) error {
		err := s.secretManager.UpdateSecret(ctx, connection.ConnectionID, data)
		if err != nil {
			return err
		}

		err = s.repo.Create(ctx, connection, session)

		if err != nil {
			err = s.secretManager.DeleteSecret(ctx, connection.ConnectionID)
			if err != nil {
				// rollback on failure
				return fmt.Errorf("unable to rollback connectionID failed to create reference : %s", connection.ConnectionID)
			}

		}
		return err
	})
	return err
}

func (s *ConnectionService) UpdateConnection(ctx context.Context, connection *connections.Connection, data *connections.ConnectionData) error {
	err := s.repo.T().DoTransaction(ctx, func(session *xorm.Session) error {
		err := s.secretManager.UpdateSecret(ctx, connection.ConnectionID, data)
		if err != nil {
			return err
		}

		err = s.repo.Update(ctx, connection, session)

		if err != nil {
			err = s.secretManager.DeleteSecret(ctx, connection.ConnectionID)
			if err != nil {
				// rollback on failure
				return fmt.Errorf("unable to rollback connectionID failed to update reference : %s", connection.ConnectionID)
			}

		}
		return err
	})
	return err
}
