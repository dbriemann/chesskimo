package base

// Piece, Color, Square and Info are basically the same type (aliases)
// but for clear declarations we use different type names.
type Piece uint8
type Color = Piece
type Square = Piece
type Info = Piece

const (
	// Various
	NONE uint8 = 0
	// OTB is off the board and used as a non-index in piece lists.
	OTB Square = 0x7F

	// Colors:
	// BLACK and WHITE are used for pieces, DARK and LIGHT are used for squares.
	BLACK           Color = 0   // 00000000
	WHITE           Color = 1   // 00000001
	COLOR_ONLY_MASK Color = 1   // 00000001
	COLOR_TEST_MASK Color = 129 // 10000001

	DARK  Color = iota // 0
	LIGHT              // 1

	// Pieces:
	// TODO: should NO_PIECE equal EMPTY??
	NO_PIECE   Piece = 0   // For special usage only
	EMPTY      Piece = 128 // 10000000
	PAWN       Piece = 2   // 00000010
	KNIGHT     Piece = 4   // 00000100
	BISHOP     Piece = 8   // 00000000
	ROOK       Piece = 16  // 00010000
	QUEEN      Piece = 32  // 00100000
	KING       Piece = 64  // 01000000
	PIECE_MASK Piece = 126 // 01111110
)

const (
	// occupancy types by color and piece
	BPAWN   Piece = PAWN | BLACK
	BKNIGHT Piece = KNIGHT | BLACK
	BBISHOP Piece = BISHOP | BLACK
	BROOK   Piece = ROOK | BLACK
	BQUEEN  Piece = QUEEN | BLACK
	BKING   Piece = KING | BLACK
	WPAWN   Piece = PAWN | WHITE
	WKNIGHT Piece = KNIGHT | WHITE
	WBISHOP Piece = BISHOP | WHITE
	WROOK   Piece = ROOK | WHITE
	WQUEEN  Piece = QUEEN | WHITE
	WKING   Piece = KING | WHITE
)

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

const (
	INFO_NONE Info = iota
	INFO_ATTACKED
	INFO_MAYBE_PINNED
	INFO_PINNED
)

var (
	// BLACK == 0, WHITE == 1
	PAWN_PUSH_DIRS    = [2]int8{DOWN, UP}
	PAWN_CAPTURE_DIRS = [2][2]int8{{DOWN_LEFT, DOWN_RIGHT}, {UP_LEFT, UP_RIGHT}}
	PAWN_BASE_RANK    = [2]Square{6, 1}
	PAWN_PROMOTE_RANK = [2]Square{0, 7}

	KNIGHT_DIRS         = [8]int8{UP + UP_LEFT, UP + UP_RIGHT, DOWN + DOWN_LEFT, DOWN + DOWN_RIGHT, LEFT + UP_LEFT, LEFT + DOWN_LEFT, RIGHT + UP_RIGHT, RIGHT + DOWN_RIGHT}
	DIAGONAL_DIRS       = [4]int8{UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}
	ORTHOGONAL_DIRS     = [4]int8{LEFT, UP, RIGHT, DOWN}
	KING_DIRS           = [8]int8{RIGHT, LEFT, UP, DOWN, UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}
	CASTLING_PATH_SHORT = [2][2]Square{{0x75, 0x76}, {0x5, 0x6}}
	CASTLING_PATH_LONG  = [2][3]Square{{0x73, 0x72, 0x71}, {0x1, 0x2, 0x3}}

	INFO_BOARD_INDEXES = [64]Square{
		0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
		0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
		0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
		0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
		0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f,
	}
)

var (
	PrintMap = map[Piece]string{
		BPAWN:   "p",
		BKNIGHT: "n",
		BBISHOP: "b",
		BROOK:   "r",
		BQUEEN:  "q",
		BKING:   "k",
		WPAWN:   "P",
		WKNIGHT: "N",
		WBISHOP: "B",
		WROOK:   "R",
		WQUEEN:  "Q",
		WKING:   "K",
		EMPTY:   ".",
	}

	PrintBoardIndex = map[Square]string{
		0x70: "a8", 0x71: "b8", 0x72: "c8", 0x73: "d8", 0x74: "e8", 0x75: "f8", 0x76: "g8", 0x77: "h8",
		0x60: "a7", 0x61: "b7", 0x62: "c7", 0x63: "d7", 0x64: "e7", 0x65: "f7", 0x66: "g7", 0x67: "h7",
		0x50: "a6", 0x51: "b6", 0x52: "c6", 0x53: "d6", 0x54: "e6", 0x55: "f6", 0x56: "g6", 0x57: "h6",
		0x40: "a5", 0x41: "b5", 0x42: "c5", 0x43: "d5", 0x44: "e5", 0x45: "f5", 0x46: "g5", 0x47: "h5",
		0x30: "a4", 0x31: "b4", 0x32: "c4", 0x33: "d4", 0x34: "e4", 0x35: "f4", 0x36: "g4", 0x37: "h4",
		0x20: "a3", 0x21: "b3", 0x22: "c3", 0x23: "d3", 0x24: "e3", 0x25: "f3", 0x26: "g3", 0x27: "h3",
		0x10: "a2", 0x11: "b2", 0x12: "c2", 0x13: "d2", 0x14: "e2", 0x15: "f2", 0x16: "g2", 0x17: "h2",
		0x00: "a1", 0x01: "b1", 0x02: "c1", 0x03: "d1", 0x04: "e1", 0x05: "f1", 0x06: "g1", 0x07: "h1",
	}
)

func (c Color) FlipColor() Color {
	// c MUST BE 0 or 1
	return c ^ 1
}

func (p Piece) PieceColor() Color {
	return p & COLOR_ONLY_MASK
}

func (p Piece) HasColor(color Color) bool {
	return (COLOR_TEST_MASK & p) == color
}

func (p Piece) IsType(piece Piece) bool {
	return (p & piece) != 0
}

func (sq Square) IsEmpty() bool {
	return sq == EMPTY
}

func (sq Square) IsLegal() bool {
	return (sq & 0x88) == 0
}

func (sq Square) SquareColor() Color {
	// Dark squares have an even index, light squares have an odd one.
	return sq & 1
}

func (sq Square) Rank() Square {
	return sq >> 4
}

func (sq Square) File() Square {
	return sq & 7
}

func (sq Square) IsPawnBaseRank(color Color) bool {
	// color MUST BE 0 or 1
	return PAWN_BASE_RANK[color] == sq.Rank()
}

func (sq Square) IsPawnPromoting(color Color) bool {
	// color MUST BE 0 or 1
	return PAWN_PROMOTE_RANK[color] == sq.Rank()
}

func (sq Square) ToInfoIndex() Square {
	return sq + 8
}
