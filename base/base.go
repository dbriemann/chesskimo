package base

type Color uint8
type Piece uint8
type Square uint8

const (
	// various
	NONE = 0
)

const (
	// colors
	BLACK Color = iota // 0
	WHITE              // 1
)

const (
	// piece types
	EMPTY  Piece = iota // 0
	PAWN                // 1
	KNIGHT              // 2
	BISHOP              // 3
	ROOK                // 4
	QUEEN               // 5
	KING                // 6
)

const (
	// occupancy types by color and piece
	BPAWN   = PAWN
	BKNIGHT = KNIGHT
	BBISHOP = BISHOP
	BROOK   = ROOK
	BQUEEN  = QUEEN
	BKING   = KING
	WPAWN   = PAWN | Piece(WHITE<<7)
	WKNIGHT = KNIGHT | Piece(WHITE<<7)
	WBISHOP = BISHOP | Piece(WHITE<<7)
	WROOK   = ROOK | Piece(WHITE<<7)
	WQUEEN  = QUEEN | Piece(WHITE<<7)
	WKING   = KING | Piece(WHITE<<7)
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
)

func (sq Square) isLegal() bool {
	return (sq & 0x88) == 0
}

func (sq Square) color() Color {
	// Black squares have an even index, White squares have an odd one.
	return Color(sq) & 1
}
