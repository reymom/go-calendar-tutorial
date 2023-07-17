//go:build mocking_app

package config

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
)

const ConfigJsonName = "calendar.json"

var (
	Version   = "UNKNOWN"
	BuildDate = "UNKNOWN"
)

type Config struct {
	ConnectionStringRead  string `json:"connectionStringRead"`
	ConnectionStringWrite string `json:"connectionStringWrite"`
	NumberOfTasks         uint   `json:"numOfTasks"`
	DaysInRange           uint   `json:"daysInRange"`
}

func (c *Config) Validate() error {
	if c.NumberOfTasks < 1 {
		log.Warn().Msgf("Number of tasks set to zero")
	}
	return nil
}

func GenerateConfig(path string) (*Config, error) {
	log.Info().Msgf(" ---------- Generating Config from path %s -------------", path)
	b, e := os.ReadFile(path)
	if e != nil {
		return nil, e
	}
	var config Config
	e = json.Unmarshal(b, &config)
	if e != nil {
		return nil, e
	}
	e = config.Validate()
	if e != nil {
		return nil, e
	}
	return &config, nil
}
