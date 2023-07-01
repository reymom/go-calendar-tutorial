package model

const (
	PriorityTypeIdLow PriorityTypeId = iota
	PriorityTypeIdMiddle
	PriorityTypeIdHigh
	lastPriorityTypeId
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
	lastColorId
)

const (
	timeScaleIdDay TimeScaleId = iota
	timeScaleIdWeek
	timeScaleIdMonth
	timeScaleIdYear
	lastTimeScaleId
)

const (
	lowerYearBond uint = 2022
	upperWeekBond uint = 53
)
