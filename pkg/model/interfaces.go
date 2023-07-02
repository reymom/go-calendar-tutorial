package model

import (
	"context"
	"time"
)

type TasksDao interface {
	//CreateTask creates the new task derived from AddableTask.
	//Returns the Task, with the fields of AddableTask together with the automatic ones created at initialization.
	//Returns an error if the Task could not be added
	CreateTask(ctx context.Context, task *AddableTask) (*Task, error)
	//RemoveTask removes the task with taskId.
	//Returns an error if no Task with taskId exists.
	RemoveTask(ctx context.Context, taskId string) error
	//EditTask edits the task with the taskId.
	//Returns the udpated Task.
	//Returns the error if no Task with taskId exists.
	EditTask(ctx context.Context, taskId string, task *AddableTask) (*Task, error)
	//SetCompleted edits the field Task.Completed and sets it to true.
	//Returns the error if no Task with taskId exists.
	SetCompleted(ctx context.Context, taskId string) error
	//ListTasks returns the Task list for the tasks that fall inside the time range defined in TimeFilter.
	ListTasks(ctx context.Context, filter *TimeFilter) ([]Task, error)
}

type TimeFilter interface {
	Validate() error
	GetTimeScaleId() TimeScaleId
	GetTimeRange() [2]time.Time
}
