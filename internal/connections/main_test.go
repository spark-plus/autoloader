package connections

import (
	"os"
	"testing"

	"github.com/spark-plus/autoloader/testutils"
)

func TestMain(m *testing.M) {
	testutils.StartMySQLContainer(m)
	defer testutils.StopMySQLContainer(m)

	code := m.Run()

	os.Exit(code)
}
