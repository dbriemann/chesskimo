// Package base provides basic constants and functions valid for
// all types of bitboard representations.
package base

type Color uint8
type Piece int
type Square uint8
type BB uint64

const (
	// various
	NONE = 0

	// files
	FILE_A = 0x0101010101010101
	FILE_B = 0x0202020202020202
	FILE_C = 0x0404040404040404
	FILE_D = 0x0808080808080808
	FILE_E = 0x1010101010101010
	FILE_F = 0x2020202020202020
	FILE_G = 0x4040404040404040
	FILE_H = 0x8080808080808080

	// ranks
	RANK_1 = 0x00000000000000FF
	RANK_2 = 0x000000000000FF00
	RANK_3 = 0x0000000000FF0000
	RANK_4 = 0x00000000FF000000
	RANK_5 = 0x000000FF00000000
	RANK_6 = 0x0000FF0000000000
	RANK_7 = 0x00FF000000000000
	RANK_8 = 0xFF00000000000000

	// bitmask for both en passent ranks
	EP_RANKS = RANK_3 | RANK_6 //0x0000FF0000FF0000
)
const (
	// colors
	WHITE Color = iota // 0
	BLACK              // 1
)

const (
	// piece types
	PAWN   Piece = iota // 0
	KNIGHT              // 1
	BISHOP              // 2
	ROOK                // 3
	QUEEN               // 4
	KING                // 5
)

const (
	MASK_PIECE = 0x7
	MASK_COLOR = 0x8
)

const (
	// occupancy types by color and piece
	WPAWN   Piece = iota // 0
	WKNIGHT              // 1
	WBISHOP              // 2
	WROOK                // 3
	WQUEEN               // 4
	WKING                // 5
	EMPTY                // 6
	_                    // 7 - skip
	BPAWN                // 8
	BKNIGHT              // 9
	BBISHOP              // 10
	BROOK                // 11
	BQUEEN               // 12
	BKING                // 13
)

func IdxBitmask(i BB) BB {
	return 1 << i
}
