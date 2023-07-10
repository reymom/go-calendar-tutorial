package mapping

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
)

type PgTimeRange struct {
	StartsAt   pgtype.Timestamptz
	FinishesAt pgtype.Timestamptz
}

func PgTimeRangeFromTimeFilter(filter model.TimeFilter) (*PgTimeRange, error) {
	e := filter.Validate()
	if e != nil {
		return nil, e
	}

	timeRange := filter.GetTimeRange()

	var out = &PgTimeRange{}
	e = out.setStartsAt(timeRange[0])
	if e != nil {
		return nil, e
	}
	e = out.setFinishesAt(timeRange[1])
	if e != nil {
		return nil, e
	}
	return out, nil
}

func (t *PgTimeRange) setStartsAt(startsAt time.Time) error {
	pgTime, e := pgTimestamptzFromTime(startsAt)
	if e != nil {
		return e
	}
	t.StartsAt = pgTime
	return nil
}

func (t *PgTimeRange) setFinishesAt(finishesAt time.Time) error {
	pgTime, e := pgTimestamptzFromTime(finishesAt)
	if e != nil {
		return e
	}
	t.FinishesAt = pgTime
	return nil
}
