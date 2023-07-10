package model

import "time"

type TimeScaleId uint8

func (t TimeScaleId) Validate() error {
	if t < lastTimeScaleId {
		return nil
	}
	return errUnknownTimeScaleTypeId
}

func NewYearlyFilter(year uint) *YearlyFilter {
	return &YearlyFilter{
		year: year,
	}
}

type YearlyFilter struct {
	year uint
}

func (y *YearlyFilter) Validate() error {
	e := y.getTimeScaleId().Validate()
	if e != nil {
		return e
	}
	if y.year <= lowerYearBond {
		return errInvalidYear
	}
	return nil
}

func (y *YearlyFilter) getTimeScaleId() TimeScaleId {
	return timeScaleIdYear
}

func (y *YearlyFilter) GetTimeRange() [2]time.Time {
	date := time.Date(int(y.year), time.January, 0, 0, 0, 0, 0, y.getLocation())
	return [...]time.Time{date, date.AddDate(1, 0, 0)}
}

func (y *YearlyFilter) getLocation() *time.Location {
	return time.UTC
}

type MonthlyFilter struct {
	year  *YearlyFilter
	month time.Month
}

func NewMonthlyFilter(month time.Month, year uint) *MonthlyFilter {
	return &MonthlyFilter{
		year:  NewYearlyFilter(year),
		month: month,
	}
}

func (m *MonthlyFilter) Validate() error {
	e := m.getTimeScaleId().Validate()
	if e != nil {
		return e
	}
	return m.year.Validate()
}

func (m *MonthlyFilter) getTimeScaleId() TimeScaleId {
	return timeScaleIdMonth
}

func (m *MonthlyFilter) GetTimeRange() [2]time.Time {
	date := time.Date(int(m.year.year), m.month, 0, 0, 0, 0, 0, m.year.getLocation())
	return [...]time.Time{date, date.AddDate(0, 1, 0)}
}

type WeeklyFilter struct {
	year YearlyFilter
	week uint
}

func NewWeeklyFilter(week, year uint) *WeeklyFilter {
	return &WeeklyFilter{
		year: *NewYearlyFilter(year),
		week: week,
	}
}

func (w *WeeklyFilter) Validate() error {
	e := w.getTimeScaleId().Validate()
	if e != nil {
		return e
	}
	if w.week >= lowerYearBond {
		return errInvalidWeek
	}
	return w.year.Validate()
}

func (w *WeeklyFilter) getTimeScaleId() TimeScaleId {
	return timeScaleIdWeek
}

func (w *WeeklyFilter) GetTimeRange() [2]time.Time {
	// Start from the middle of the year:
	date := time.Date(int(w.year.year), 7, 1, 0, 0, 0, 0, w.year.getLocation())

	// Roll back to Monday:
	if wd := date.Weekday(); wd == time.Sunday {
		date = date.AddDate(0, 0, -6)
	} else {
		date = date.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, wk := date.ISOWeek()
	date = date.AddDate(0, 0, (int(w.week)-wk)*7)

	return [...]time.Time{date, date.AddDate(0, 0, 6)}
}

type DaylyFilter struct {
	month *MonthlyFilter
	day   uint
}

func NewDaylyFilter(day uint, month time.Month, year uint) *DaylyFilter {
	return &DaylyFilter{
		month: NewMonthlyFilter(month, year),
		day:   day,
	}
}

func (d *DaylyFilter) Validate() error {
	e := d.getTimeScaleId().Validate()
	if e != nil {
		return e
	}
	return d.month.year.Validate()
}

func (d *DaylyFilter) getTimeScaleId() TimeScaleId {
	return timeScaleIdDay
}

func (d *DaylyFilter) GetTimeRange() [2]time.Time {
	date := time.Date(int(d.month.year.year), d.month.month, int(d.day), 0, 0, 0, 0, d.month.year.getLocation())
	return [...]time.Time{date, date.AddDate(0, 0, 1)}
}
