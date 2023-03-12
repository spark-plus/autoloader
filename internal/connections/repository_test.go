package connections_test

import (
	"context"
	"testing"
	"time"

	"github.com/spark-plus/autoloader/internal/databases"
	"github.com/spark-plus/autoloader/testutils"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spark-plus/autoloader/internal/connections"
	"github.com/stretchr/testify/require"
	"xorm.io/xorm"
)

func NewTestDBSession(t *testing.T) *databases.DBSession {
	db, err := xorm.NewEngine("mysql", testutils.MysqlConnStr)
	require.NoError(t, err)

	// Drop existing tables (if any) and create new tables
	err = db.DropTables(&connections.Connection{})
	require.NoError(t, err)
	err = db.CreateTables(&connections.Connection{})
	require.NoError(t, err)

	// Return a session with the transaction
	return databases.NewDBSession(db)
}

func TestConnectionRepository_Create(t *testing.T) {
	ctx := context.Background()
	db := NewTestDBSession(t)
	repo := connections.NewConnectionRepository(db)
	tx, err := repo.BeginTx(ctx)

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	conn := &connections.Connection{
		ConnectionID:   "test-connection",
		ConnectionType: connections.MySQLConnType,
		ConnectionRef: connections.ConnectionRef{
			Location:     connections.SSMLocation,
			LocationPath: "/conf/run/ssm/local",
		},
		Description: "Test Connection",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = repo.CreateWTx(ctx, conn, tx)
	assert.NoError(t, err)
	assert.NotZero(t, conn.ID)
}
