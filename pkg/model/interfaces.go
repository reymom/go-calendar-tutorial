package model

import (
	"context"
	"time"
)

type TasksDao interface {
	CreateTask(ctx context.Context, task *AddableTask) (*Task, error)
	RemoveTask(ctx context.Context) (bool, error)
	EditTask(ctx context.Context, task *AddableTask) (*Task, error)
	SetCompleted(ctx context.Context) (bool, error)

	ListTasks(ctx context.Context, filter *TimeFilter) ([]Task, error)
}

type TimeFilter interface {
	Validate() error
	GetTimeScaleId() TimeScaleId
	GetTimeRange() [2]time.Time
}
