package connections

import (
	"context"
	"fmt"

	"github.com/sparkster/autoloader/internal/databases"
	"xorm.io/xorm"
)

// ConnectionRepository is a repository for managing connections in RDBMS
type ConnectionRepository struct {
	db *databases.DBSession
}

// NewConnectionRepository creates a new ConnectionRepository instance
func NewConnectionRepository(db *databases.DBSession) *ConnectionRepository {
	return &ConnectionRepository{db: db}
}

// CreateConnection creates a new connection reference in the RDBMS
func (r *ConnectionRepository) Create(ctx context.Context, connection *Connection) error {
	txn, err := r.db.BeginTx(ctx)
	if err != nil {
		return err
	}

	session := txn.(*databases.ConnectionTx).Session

	defer session.Close()

	if _, err := session.Insert(connection); err != nil {
		session.Rollback()
		return err
	}

	if err := session.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *ConnectionRepository) CreateWTx(ctx context.Context, session *xorm.Session, connection *Connection) error {
	_, err := session.Insert(connection)
	if err != nil {
		return fmt.Errorf("failed to insert connection: %v", err)
	}
	return nil
}

// UpdateConnection updates an existing connection reference in the RDBMS
func (r *ConnectionRepository) UpdateWTx(ctx context.Context, session *xorm.Session, conn *Connection) error {
	if _, err := session.ID(conn.ID).Update(conn); err != nil {
		return fmt.Errorf("failed to update connection: %v", err)
	}
	return nil
}

// DeleteConnection deletes an existing connection reference in the RDBMS
func (r *ConnectionRepository) DeleteWT(ctx context.Context, session *xorm.Session, conn *Connection) error {
	if _, err := session.ID(conn.ID).Delete(conn); err != nil {
		return fmt.Errorf("failed to delete connection: %v", err)
	}
	return nil
}

// GetConnectionByID returns the connection reference with the given ID
func (r *ConnectionRepository) GetByID(ctx context.Context, id int64) (*Connection, error) {
	txn, err := r.db.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	conn := &Connection{}

	session := txn.(*databases.ConnectionTx).Session

	has, err := session.ID(id).Get(conn)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("Connection with ID %d not found", id)
	}
	return conn, nil

}

// GetConnectionByID returns the connection reference with the given ID
func (r *ConnectionRepository) GetByConnectionId(ctx context.Context, id string) (*Connection, error) {
	txn, err := r.db.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	conn := &Connection{}

	session := txn.(*databases.ConnectionTx).Session

	has, err := session.Where("connection_id = ?", id).Get(conn)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("Connection with ConnectionID %s not found", id)
	}
	return conn, nil

}
