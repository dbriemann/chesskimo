package base

type Color uint8
type Piece uint8
type Square uint8

const (
	// various
	NONE = 0

	// OTB is off the board and used as a non-index in piece lists.
	OTB Square = 0x7F
)

const (
	// colors:
	// BLACK and WHITE are used for pieces, DARK and LIGHT are used for squares.
	BLACK                Color = 64  // 01000000 -> second highest bit in uint8
	WHITE                Color = 128 // 10000000 -> highest bit in uint8
	COLOR_MASK           Color = 192 // 11000000 -> bith highest bits in uint8
	BLACK_INDEX                = 0
	WHITE_INDEX                = 1
	WHITE_OR_BLACK_SHIFT       = 7

	DARK  Color = iota // 0
	LIGHT              // 1
)

const (
	// piece types
	EMPTY      Piece = 0  // 0
	PAWN       Piece = 1  // 00000001
	KNIGHT     Piece = 2  // 00000010
	BISHOP     Piece = 4  // 00000100
	ROOK       Piece = 8  // 00001000
	QUEEN      Piece = 16 // 00010000
	KING       Piece = 32 // 00100000
	PIECE_MASK Piece = 63 // 00111111
)

const (
	// occupancy types by color and piece
	BPAWN   Piece = PAWN | Piece(BLACK)
	BKNIGHT Piece = KNIGHT | Piece(BLACK)
	BBISHOP Piece = BISHOP | Piece(BLACK)
	BROOK   Piece = ROOK | Piece(BLACK)
	BQUEEN  Piece = QUEEN | Piece(BLACK)
	BKING   Piece = KING | Piece(BLACK)
	WPAWN   Piece = PAWN | Piece(WHITE)
	WKNIGHT Piece = KNIGHT | Piece(WHITE)
	WBISHOP Piece = BISHOP | Piece(WHITE)
	WROOK   Piece = ROOK | Piece(WHITE)
	WQUEEN  Piece = QUEEN | Piece(WHITE)
	WKING   Piece = KING | Piece(WHITE)
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

func (c Color) ToIndex() Color {
	return c >> WHITE_OR_BLACK_SHIFT
}

func (c Color) FlipColor() Color {
	//TODO
	return 0
}

func (p Piece) Color() Color {
	return Color(p) & COLOR_MASK
}

func (p Piece) HasColor(color Color) bool {
	return (color & Color(p)) != 0
}

func (sq Square) IsLegal() bool {
	return (sq & 0x88) == 0
}

func (sq Square) Color() Color {
	// Dark squares have an even index, light squares have an odd one.
	return Color(sq) & 1
}

func (sq Square) Rank() Square {
	return sq >> 4
}

func (sq Square) File() Square {
	return sq & 7
}

func (sq Square) IsPawnBaseRank(colorIDX Color) bool {
	return PAWN_BASE_RANK[colorIDX] == sq.Rank()
}

func (sq Square) IsPawnPromoting(colorIDX Color) bool {
	return PAWN_PROMOTE_RANK[colorIDX] == sq.Rank()
}
