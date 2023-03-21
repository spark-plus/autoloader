package connections

import (
	"context"
	"fmt"

	"github.com/spark-plus/autoloader/internal/databases"
	"xorm.io/xorm"
)

type ConnectionRepo interface {
	T() databases.Tx
	Create(ctx context.Context, conn *Connection, session *xorm.Session) error
	Update(ctx context.Context, conn *Connection, session *xorm.Session) error
	Delete(ctx context.Context, conn *Connection, session *xorm.Session) error
	GetByID(ctx context.Context, id int64) (*Connection, error)
	GetByConnectionId(ctx context.Context, id string) (*Connection, error)
}

// connectionRepository is a repository for managing connections in RDBMS
type connectionRepository struct {
	tx databases.Tx
}

// NewconnectionRepository creates a new connectionRepository instance
func NewconnectionRepository(engine *xorm.Engine) ConnectionRepo {
	return &connectionRepository{tx: databases.NewTransactionManager(engine)}
}

func (r *connectionRepository) T() databases.Tx {
	return r.tx
}

func (r *connectionRepository) Create(ctx context.Context, connection *Connection, session *xorm.Session) error {
	_, err := session.Insert(connection)
	if err != nil {
		return fmt.Errorf("failed to insert connection: %v", err)
	}
	return nil
}

// UpdateConnection updates an existing connection reference in the RDBMS
func (r *connectionRepository) Update(ctx context.Context, conn *Connection, session *xorm.Session) error {
	if conn.ID == 0 {
		return fmt.Errorf("cannot update non existent Id: %d", conn.ID)
	}
	if _, err := session.ID(conn.ID).Update(conn); err != nil {
		return fmt.Errorf("failed to update connection: %v", err)
	}
	return nil
}

// DeleteConnection deletes an existing connection reference in the RDBMS
func (r *connectionRepository) Delete(ctx context.Context, conn *Connection, session *xorm.Session) error {
	if _, err := session.ID(conn.ID).Delete(conn); err != nil {
		return fmt.Errorf("failed to delete connection: %v", err)
	}
	return nil
}

// GetConnectionByID returns the connection reference with the given ID
func (r *connectionRepository) GetByID(ctx context.Context, id int64) (*Connection, error) {
	session := r.tx.NewSession()
	conn := &Connection{}

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
func (r *connectionRepository) GetByConnectionId(ctx context.Context, id string) (*Connection, error) {
	session := r.tx.NewSession()
	conn := &Connection{}

	has, err := session.Where("connection_id = ?", id).Get(conn)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("Connection with ConnectionID %s not found", id)
	}
	return conn, nil

}
