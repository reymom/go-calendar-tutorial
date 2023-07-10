package model

const (
	PriorityTypeIdLow PriorityTypeId = iota
	PriorityTypeIdMiddle
	PriorityTypeIdHigh
	LastPriorityTypeId
)

const (
	ColorIdRed ColorId = iota
	ColorIdYellow
	ColorIdBlue
	ColorIdOrange
	ColorIdGreen
	ColorIdViolet
	ColorIdCyan
	ColorIdBlack
	ColorIdWhite
	LastColorId
)

const (
	timeScaleIdDay TimeScaleId = iota
	timeScaleIdWeek
	timeScaleIdMonth
	timeScaleIdYear
	lastTimeScaleId
)

const (
	lowerYearBond uint = 2021
	upperWeekBond uint = 53
)
