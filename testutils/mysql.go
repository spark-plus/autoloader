package testutils

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	mysqlC       testcontainers.Container
	MysqlConnStr string
)

func StartMySQLContainer(t *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_DATABASE":      "testdb",
			"MYSQL_USER":          "user",
			"MYSQL_PASSWORD":      "password",
			"MYSQL_ROOT_PASSWORD": "rootpassword",
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("port: 3306  MySQL Community Server - GPL"),
			wait.ForListeningPort("3306/tcp"),
		),
		Name: "spark-plus-mysql",
	}
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            true,
	})
	if err != nil {
		log.Fatalf("failed to start MySQL container: %v", err)
	}

	ip, err := mysqlC.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get MySQL container host: %v", err)
	}

	port, err := mysqlC.MappedPort(ctx, "3306")
	if err != nil {
		log.Fatalf("failed to get MySQL container port: %v", err)
	}

	MysqlConnStr = fmt.Sprintf("user:password@tcp(%s:%s)/testdb", ip, port.Port())

	log.Printf("MySQL container started with connection string: %s", MysqlConnStr)
}

func StopMySQLContainer(t *testing.M) {
	ctx := context.Background()

	err := mysqlC.Terminate(ctx)
	if err != nil {
		log.Fatalf("failed to stop MySQL container: %v", err)
	}

	log.Println("MySQL container stopped")
}
