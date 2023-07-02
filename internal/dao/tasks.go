package dao

import (
	"context"

	"github.com/reymom/go-calendar-tutorial/internal/dao/mapping"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
)

func (p *PsqlDao) CreateTask(ctx context.Context, task *model.AddableTask) (*model.Task, error) {
	const query = "INSERT INTO Tasks(name, description, starts_at, finished_at, priority, color) " +
		"VALUES ($1, $2, $3, $4) " +
		"Returning Users.display_id, Users.completed"

	in, e := mapping.PgAddableTaskFromAddableTask(task)
	if e != nil {
		return nil, e
	}
	var out = mapping.PgTask{
		PgAddableTask: *in,
	}
	e = p.writePool.QueryRow(ctx, query, &in.Name, &in.Description, &in.StartsAt, &in.FinishesAt, &in.Priority, &in.Color).
		Scan(&out.TaskId, &out.Completed)
	if e != nil {
		return nil, e
	}

	return out.Parse()
}

func (p *PsqlDao) RemoveTask(ctx context.Context, taskId string) error {
	panic("hello")
}

func (p *PsqlDao) EditTask(ctx context.Context, taskId string, task *model.AddableTask) (*model.Task, error) {
	panic("hello")
}

func (p *PsqlDao) SetCompleted(ctx context.Context, taskId string) error {
	panic("hello")
}

func (p *PsqlDao) ListTasks(ctx context.Context, filter *model.TimeFilter) ([]model.Task, error) {
	panic("hello")
}
