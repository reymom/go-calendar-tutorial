//go:build mocking_app && unit_test

package config

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	connectionString = "postgresql://calendar_dev_user:calendar_dev_password@172.17.0.2:5432/calendar_dev_user"
)

var configPath = path.Join("testdata", ConfigJsonName)

func TestConfig(t *testing.T) {
	a := assert.New(t)

	exampleConfig := Config{
		ConnectionStringRead:  connectionString,
		ConnectionStringWrite: connectionString,
		NumberOfTasks:         3,
		DaysInRange:           7,
	}

	b, e := json.MarshalIndent(exampleConfig, "", "  ")
	a.NoError(e)
	configJsonFile, e := os.Create(configPath)
	a.NoError(e)
	_, e = configJsonFile.Write(b)
	a.NoError(e)
	e = configJsonFile.Close()
	a.NoError(e)

	_, e = GenerateConfig(configPath)
	a.NoError(e)
}
