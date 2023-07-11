//go:build mocking_app

package main

import (
	"flag"
	"os"
	"path"

	"github.com/reymom/go-calendar-tutorial/cmd/mocking/config"
	"github.com/reymom/go-calendar-tutorial/internal/dao"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msgf(" ---------- Mocking App, Version %s, Build date %s -------------", config.Version, config.BuildDate)
	configFilePath := flag.String("c", path.Join("conf/"+config.ConfigJsonName), "config file path")
	flag.Parse()
	conf, e := config.GenerateConfig(*configFilePath)
	if e != nil {
		log.Err(e).Msgf("Error while generating config")
		os.Exit(1)
	}

	mockingApp, e := newMockingApp(conf)
	if e != nil {
		log.Err(e).Msgf("Error while initializing mocking app")
		os.Exit(1)
	}
	e = mockingApp.mockData(conf)
	if e != nil {
		log.Err(e).Msgf("Error while mocking data")
		os.Exit(1)
	}

}

type mockingApp struct {
	tasksDao      *dao.PsqlDao
	numberOfTasks uint
	daysInRange   uint
}

func newMockingApp(conf *config.Config) (*mockingApp, error) {
	if conf == nil {
		return nil, model.ErrNilNotAllowed
	}
	taskDao, e := dao.NewPsqlDao(
		&dao.Config{
			ConnectionStringRead:  conf.ConnectionStringRead,
			ConnectionStringWrite: conf.ConnectionStringWrite,
			WriteEnabled:          true,
			MaxReadConnections:    2,
			MaxWriteConnections:   2,
		},
	)
	if e != nil {
		return nil, e
	}
	taskDao.MockingCleanUp()
	return &mockingApp{
		tasksDao:      taskDao,
		numberOfTasks: conf.NumberOfTasks,
		daysInRange:   conf.DaysInRange,
	}, nil
}

func (m *mockingApp) mockData(conf *config.Config) error {
	for i := 0; i < int(m.numberOfTasks); i++ {
		e := m.tasksDao.MockTask(conf.DaysInRange)
		if e != nil {
			return e
		}
	}
	return nil
}
