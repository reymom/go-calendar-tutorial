package model

type ErrorHandler interface {
	HandleError(e error)
}

type ConstError string

func (c ConstError) Error() string {
	return string(c)
}

// psql-related errors
const (
	ErrNilNotAllowed             ConstError = "nil not allowed"
	ErrFakeWritePool             ConstError = "no connection supported for fake write pool"
	ErrReadConnectionStringEmpty ConstError = "empty connection string"
	ErrTaskNotFound              ConstError = "task not found"
)

// mapping errors
const (
	ErrTextToLong             ConstError = "Text too long"
	ErrOutOfBoundsOfPGInt2    ConstError = "Int2 out of bonds"
	ErrInt2NotCastableToUint8 ConstError = "Int2 out of bounds for uint8"
	ErrTimestamptzNotPresent  ConstError = "Timestamptz not present"
	ErrTextNotPresent         ConstError = "Text not present"
	ErrUuidNotPresent         ConstError = "Uuid not present"
	ErrInt2NotPresent         ConstError = "Int2 not present"
	ErrBoolNotPresent         ConstError = "Bool not present"
)

// model types error
const (
	errUnknownPriorityTypeId  ConstError = "unknown priority type"
	errUnknownColorId         ConstError = "unknown color"
	errUnknownTimeScaleTypeId ConstError = "unknown time scale type"
	errInvalidYear            ConstError = "invalid year"
	errInvalidWeek            ConstError = "invalid week"
)
