//go:build integration_test

package integrationtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testEnv *testEnvironment

const connectionString = "postgresql://calendar_dev_user:calendar_dev_password@localhost:5432/calendar_dev_user"

func TestSetup(t *testing.T) {
	t.Run("Setup Dao Connection", daoTest)
	t.Run("Task Dao", taskDaoTests)
}

func daoTest(t *testing.T) {
	a := assert.New(t)
	var e error
	testEnv, e = newTestEnvironment(connectionString, connectionString, t)
	a.NoError(e)
	t.Cleanup(testEnv.cleanup)
}
