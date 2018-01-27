package base

const (
	UP         = 16
	DOWN       = -16
	LEFT       = -1
	RIGHT      = +1
	UP_LEFT    = UP + LEFT
	UP_RIGHT   = UP + RIGHT
	DOWN_LEFT  = DOWN + LEFT
	DOWN_RIGHT = DOWN + RIGHT
)

var (
	// BLACK == 0, WHITE == 1
	PAWN_PUSH_DIRS    = [2]int8{DOWN, UP}
	PAWN_CAPTURE_DIRS = [2][2]int8{{DOWN_LEFT, DOWN_RIGHT}, {UP_LEFT, UP_RIGHT}}
	PAWN_BASE_RANK    = [2]uint8{7, 2}
)
