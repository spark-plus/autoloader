package connections_test

import (
	"context"
	"testing"

	"github.com/spark-plus/autoloader/testutils"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spark-plus/autoloader/internal/connections"
	"github.com/stretchr/testify/require"
	"xorm.io/xorm"
)

func NewTestDBSession(t *testing.T) *xorm.Engine {
	db, err := xorm.NewEngine("mysql", testutils.MysqlConnStr)
	require.NoError(t, err)

	// Drop existing tables (if any) and create new tables
	err = db.DropTables(&connections.Connection{})
	require.NoError(t, err)
	err = db.CreateTables(&connections.Connection{})
	require.NoError(t, err)

	// Return a session with the transaction
	return db
}

func TestConnectionRepository_Create(t *testing.T) {
	ctx := context.Background()
	db := NewTestDBSession(t)
	repo := connections.NewconnectionRepository(db)
	repo.T().DoTransaction(ctx, func(session *xorm.Session) error {
		conn := connections.NewConnection(
			"test-connection",
			connections.MySQLConnType,
			connections.ConnectionRef{
				Location:     connections.SSMLocation,
				LocationPath: "/conf/run/ssm/local",
			},
			"Test Connection",
		)
		err := repo.Create(ctx, conn, session)
		assert.NoError(t, err)
		assert.NotZero(t, conn.ID)
		return err
	})
}
