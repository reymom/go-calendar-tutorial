package mapping

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
)

const (
	taskNameLength        = 126
	taskDescriptionLength = 255
)

type PgTask struct {
	TaskId    pgtype.UUID
	Completed pgtype.Bool
	PgAddableTask
}

func (t *PgTask) Parse() (*model.Task, error) {
	taskId, e := t.getTaskId()
	if e != nil {
		return nil, e
	}
	name, e := t.getName()
	if e != nil {
		return nil, e
	}
	description, e := t.getDescription()
	if e != nil {
		return nil, e
	}
	startsAt, e := t.getStartsAt()
	if e != nil {
		return nil, e
	}
	finishesAt, e := t.getFinishesAt()
	if e != nil {
		return nil, e
	}
	priority, e := t.getPriorityTypeId()
	if e != nil {
		return nil, e
	}
	color, e := t.getColorId()
	if e != nil {
		return nil, e
	}
	completed, e := t.getCompleted()
	if e != nil {
		return nil, e
	}

	return &model.Task{
		TaskId:    taskId,
		Completed: completed,
		AddableTask: model.AddableTask{
			Name:        name,
			Description: description,
			StartsAt:    startsAt,
			FinishesAt:  finishesAt,
			Priority:    priority,
			Color:       color,
		},
	}, nil
}

func (t *PgTask) getTaskId() (string, error) {
	return uuidFromPGUuid(t.TaskId)
}

func (t *PgTask) getName() (string, error) {
	return stringFromPGVarchar(t.Name)
}

func (t *PgTask) getDescription() (string, error) {
	return stringFromPGVarchar(t.Description)
}

func (t *PgTask) getStartsAt() (time.Time, error) {
	return timeFromPgTimestamptz(t.StartsAt)
}

func (t *PgTask) getFinishesAt() (time.Time, error) {
	return timeFromPgTimestamptz(t.FinishesAt)
}

func (t *PgTask) getPriorityTypeId() (model.PriorityTypeId, error) {
	return uint8FromPgInt2[model.PriorityTypeId](t.Priority)
}

func (t *PgTask) getColorId() (model.ColorId, error) {
	return uint8FromPgInt2[model.ColorId](t.Color)
}

func (t *PgTask) getCompleted() (bool, error) {
	return boolFromPgBool(t.Completed)
}

type PgAddableTask struct {
	Name        pgtype.Text
	Description pgtype.Text
	StartsAt    pgtype.Timestamptz
	FinishesAt  pgtype.Timestamptz
	Priority    pgtype.Int2
	Color       pgtype.Int2
}

func PgAddableTaskFromAddableTask(task *model.AddableTask) (*PgAddableTask, error) {
	var out = &PgAddableTask{}
	e := out.setName(task.Name)
	if e != nil {
		return nil, e
	}
	e = out.setDescription(task.Description)
	if e != nil {
		return nil, e
	}
	e = out.setStartsAt(task.StartsAt)
	if e != nil {
		return nil, e
	}
	e = out.setFinishesAt(task.FinishesAt)
	if e != nil {
		return nil, e
	}
	e = out.setPriority(task.Priority)
	if e != nil {
		return nil, e
	}
	e = out.setColor(task.Color)
	if e != nil {
		return nil, e
	}
	return out, nil
}

func (a *PgAddableTask) setName(name string) error {
	pgName, e := pgTextFromString(name, taskNameLength)
	if e != nil {
		return e
	}
	a.Name = pgName
	return nil
}

func (a *PgAddableTask) setDescription(description string) error {
	pgDescription, e := pgTextFromString(description, taskDescriptionLength)
	if e != nil {
		return e
	}
	a.Description = pgDescription
	return nil
}

func (a *PgAddableTask) setStartsAt(startsAt time.Time) error {
	pgTime, e := pgTimestamptzFromTime(startsAt)
	if e != nil {
		return e
	}
	a.StartsAt = pgTime
	return nil
}

func (a *PgAddableTask) setFinishesAt(finishesAt time.Time) error {
	pgTime, e := pgTimestamptzFromTime(finishesAt)
	if e != nil {
		return e
	}
	a.FinishesAt = pgTime
	return nil
}

func (a *PgAddableTask) setPriority(priorityId model.PriorityTypeId) error {
	e := priorityId.Validate()
	if e != nil {
		return e
	}
	a.Priority, e = toPgInt2(priorityId)
	return e
}

func (a *PgAddableTask) setColor(colorId model.ColorId) error {
	e := colorId.Validate()
	if e != nil {
		return e
	}
	a.Priority, e = toPgInt2(colorId)
	return e
}
