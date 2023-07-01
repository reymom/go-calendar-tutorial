package model

type ErrorHandler interface {
	HandleError(e error)
}

type ConstError string

func (c ConstError) Error() string {
	return string(c)
}

const (
	errUnknownPriorityTypeId  ConstError = "unknown priority type"
	errUnknownColorId         ConstError = "unknown color"
	errUnknownTimeScaleTypeId ConstError = "unknown time scale type"
	errInvalidYear            ConstError = "invalid year"
	errInvalidWeek            ConstError = "invalid week"
)
