package chesskimo

// Board contains all information for a chess board state, including the board itself as 0x88 board.
//  +------------------------+
//8 |70 71 72 73 74 75 76 77 | 78 79 7a 7b 7c 7d 7e 7f
//7 |60 61 62 63 64 65 66 67 | 68 69 6a 6b 6c 6d 6e 6f
//6 |50 51 52 53 54 55 56 57 | 58 59 5a 5b 5c 5d 5e 5f
//5 |40 41 42 43 44 45 46 47 | 48 49 4a 4b 4c 4d 4e 4f
//4 |30 31 32 33 34 35 36 37 | 38 39 3a 3b 3c 3d 3e 3f
//3 |20 21 22 23 24 25 26 27 | 28 29 2a 2b 2c 2d 2e 2f
//2 |10 11 12 13 14 15 16 17 | 18 19 1a 1b 1c 1d 1e 1f
//1 | 0  1  2  3  4  5  6  7 |  8  9  a  b  c  d  e  f
//  +------------------------+
//    a  b  c  d  e  f  g  h |
// All indexes shown are in HEX. The left board represents the actual playing board, whereas there right board
// is for detection of illegal moves without heavy conditional usage.
type Board struct {
	Field       [64 * 2]Piece
	Kings       [2]Square
	CastleShort [2]bool
	CastleLong  [2]bool
	IsCheck     bool
	EpSquare    Square
	MoveNumber  uint16
	DrawCounter uint16
	Player      Color
}

// Lookup0x88 maps the indexes of a 8x8 MinBoard to the 0x88 board indexes.
var Lookup0x88 = [64]Square{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
	0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
	0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47,
	0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57,
	0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67,
	0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77,
}

func NewBoard() Board {
	b := Board{}
	return b
}

// SetStartingPosition sets the starting position of a default chess game via FEN code.
func (b *Board) SetStartingPosition() {
	err := b.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}
}

// SetFEN sets a chess position from a 'FEN' string.
func (b *Board) SetFEN(fen string) error {
	mb, err := ParseFEN(fen)
	if err != nil {
		panic(err)
	}

	// No error encountered -> set all members of the board instance.
	b.MoveNumber = mb.MoveNum
	b.DrawCounter = mb.HalfMoves
	b.EpSquare = mb.EpSquare
	b.CastleShort = mb.CastleShort
	b.CastleLong = mb.CastleLong
	b.Player = mb.Color
	for idx, piece := range mb.Squares {
		b.Field[Lookup0x88[idx]] = piece
	}

	return nil
}

func (b *Board) String() string {
	str := ""

	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			idx := 16*r + f
			str += PrintMap[b.Field[idx]]
			if f == 7 {
				str += "\n"
			}
			// Convert index to HEX.
			//			hexi := fmt.Sprintf("%x", i)
			//			trick := hexi[len(hexi)-1]
			//			if trick <= '7' {
			//				str += PrintMap[b.Field[i]]
			//			} else if trick == '8' {
			//				str += "\n"
			//			}
		}
	}

	return str
}
