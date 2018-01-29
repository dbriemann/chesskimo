package base

type Piece uint8

// Color is an alias for type Piece to clarify certain declarations
// but in the end both just represent the same things.
type Color = Piece
type Square uint8

const (
	// Various
	NONE Square = 0
	// OTB is off the board and used as a non-index in piece lists.
	OTB Square = 0x7F
)

const (
	// Colors:
	// BLACK and WHITE are used for pieces, DARK and LIGHT are used for squares.
	BLACK           Color = 0   // 00000000
	WHITE           Color = 1   // 00000001
	COLOR_ONLY_MASK Color = 1   // 00000001
	COLOR_TEST_MASK Color = 129 // 10000001

	DARK  Color = iota // 0
	LIGHT              // 1

	// Pieces:
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

func (c Color) FlipColor() Color {
	// c MUST BE 0 or 1
	return c ^ 1
}

func (p Piece) Color() Color {
	return p & COLOR_ONLY_MASK
}

func (p Piece) HasColor(color Color) bool {
	return (COLOR_TEST_MASK & p) == color
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

func (sq Square) IsPawnBaseRank(color Color) bool {
	return PAWN_BASE_RANK[color] == sq.Rank()
}

func (sq Square) IsPawnPromoting(color Color) bool {
	return PAWN_PROMOTE_RANK[color] == sq.Rank()
}
