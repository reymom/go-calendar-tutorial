package psql

import (
	"github.com/reymom/go-calendar-tutorial/internal/dao"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
)

type TaskDaoConfig struct {
	ConnectionStringRead  string
	ConnectionStringWrite string
	MaxReadConnections    uint
}

func NewTaskDao(config *TaskDaoConfig) (model.TasksDao, error) {
	internalConfig := dao.Config{
		ConnectionStringRead:  config.ConnectionStringRead,
		ConnectionStringWrite: config.ConnectionStringWrite,
		WriteEnabled:          true,
		MaxReadConnections:    config.MaxReadConnections,
	}
	return dao.NewPsqlDao(&internalConfig)
}

type ReadOnlyTaskDaoConfig struct {
	ConnectionStringRead string
	MaxReadConnections   uint
}

func NewReadOnlyTaskDao(config *ReadOnlyTaskDaoConfig) (model.TasksDao, error) {
	internalConfig := dao.Config{
		ConnectionStringRead: config.ConnectionStringRead,
		WriteEnabled:         false,
		MaxReadConnections:   config.MaxReadConnections,
	}
	return dao.NewPsqlDao(&internalConfig)
}
