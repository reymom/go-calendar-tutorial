package model

import "time"

type Task struct {
	TaskId    string
	Completed bool
	AddableTask
}

type AddableTask struct {
	Name        string
	Description string
	StartsAt    time.Time
	FinishesAt  time.Time
	Priority    PriorityTypeId
	Color       ColorId
}

type PriorityTypeId uint8

func (t PriorityTypeId) Validate() error {
	if t < LastPriorityTypeId {
		return nil
	}
	return errUnknownPriorityTypeId
}

type ColorId uint8

func (t ColorId) Validate() error {
	if t < LastColorId {
		return nil
	}
	return errUnknownColorId
}
