package dao

import (
	"context"

	"github.com/reymom/go-calendar-tutorial/internal/dao/mapping"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
)

const maxListSize uint64 = 20

func (p *PsqlDao) CreateTask(ctx context.Context, task model.AddableTask) (*model.Task, error) {
	const query = "INSERT INTO Tasks(name, description, starts_at, finishes_at, priority, color) " +
		"VALUES ($1, $2, $3, $4, $5, $6) " +
		"Returning Tasks.display_id, Tasks.completed"

	in, e := mapping.PgAddableTaskFromAddableTask(&task)
	if e != nil {
		return nil, e
	}
	var out = mapping.PgTask{
		PgAddableTask: *in,
	}
	e = p.writePool.QueryRow(ctx, query, &in.Name, &in.Description, &in.StartsAt,
		&in.FinishesAt, &in.Priority, &in.Color).Scan(&out.TaskId, &out.Completed)
	if e != nil {
		return nil, e
	}

	return out.Parse()
}

func (p *PsqlDao) EditTask(ctx context.Context, taskId string, task model.AddableTask) (*model.Task, error) {
	const query = "UPDATE Tasks " +
		"SET name=$2, description=$3, starts_at=$4, finishes_at=$5, priority=$6, color=$7 " +
		"WHERE display_id=$1 " +
		"RETURNING Tasks.name, Tasks.description, Tasks.starts_at, Tasks.finishes_at, Tasks.priority, " +
		"Tasks.color, Tasks.completed"

	in, e := mapping.PgAddableTaskFromAddableTask(&task)
	if e != nil {
		return nil, e
	}
	pgTaskId, e := mapping.PgUuidFromUuid(taskId)
	if e != nil {
		return nil, e
	}
	var out = mapping.PgTask{
		PgAddableTask: *in,
	}
	e = p.writePool.QueryRow(ctx, query, pgTaskId, &in.Name, &in.Description, &in.StartsAt,
		&in.FinishesAt, &in.Priority, &in.Color).Scan(&out.TaskId, &out.Completed)
	if e != nil {
		return nil, e
	}

	return out.Parse()
}

func (p *PsqlDao) RemoveTask(ctx context.Context, taskId string) error {
	const query = "DELETE FROM Tasks " +
		"WHERE Tasks.display_id=$1"

	pgTaskId, e := mapping.PgUuidFromUuid(taskId)
	if e != nil {
		return e
	}
	pgExec, e := p.writePool.Exec(ctx, query, pgTaskId)
	if e != nil {
		return e
	}
	if pgExec.RowsAffected() == 0 {
		return model.ErrTaskNotFound
	}
	return nil
}

func (p *PsqlDao) SetCompleted(ctx context.Context, taskId string, completed bool) error {
	const query = "UPDATE Tasks " +
		"SET completed=$2 " +
		"WHERE display_id=$1 "

	pgTaskId, e := mapping.PgUuidFromUuid(taskId)
	if e != nil {
		return e
	}
	pgCompleted, e := mapping.PgBoolFromBool(completed)
	if e != nil {
		return e
	}
	pgExec, e := p.writePool.Exec(ctx, query, pgTaskId, pgCompleted)
	if e != nil {
		return e
	}
	if pgExec.RowsAffected() == 0 {
		return model.ErrTaskNotFound
	}
	return nil
}

func (p *PsqlDao) ListTasks(ctx context.Context, filter model.TimeFilter) ([]model.Task, error) {
	pgTimeRange, e := mapping.PgTimeRangeFromTimeFilter(filter)
	if e != nil {
		return nil, e
	}

	const query = "SELECT " +
		"Tasks.display_id, Tasks.name, Tasks.description, Tasks.starts_at, " +
		"Tasks.finishes_at, Tasks.priority, Tasks.color, Tasks.completed " +
		"FROM Tasks " +
		"WHERE (starts_at) >= $1 AND (starts_at) < $2 " +
		"ORDER BY Tasks.starts_at LIMIT $3"

	rows, e := p.readPool.Query(ctx, query, &pgTimeRange.StartsAt, &pgTimeRange.FinishesAt, maxListSize)
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	ret := make([]model.Task, 0, maxListSize)
	for rows.Next() {
		var out = mapping.PgTask{}
		e := rows.Scan(&out.TaskId, &out.Name, &out.Description, &out.StartsAt,
			&out.FinishesAt, &out.Priority, &out.Color, &out.Completed)
		if e != nil {
			return nil, e
		}
		parsedOut, e := out.Parse()
		if e != nil {
			return nil, e
		}
		ret = append(ret, *parsedOut)
	}

	return ret, nil
}
