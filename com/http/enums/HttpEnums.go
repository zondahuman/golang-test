package enums

type STATUS int

const(
	SUCCESS  STATUS = iota
	FAILURE
	INIT
	EXCEPTION
	INNER_ERROR
	SYSTEM_ERROR
)
