package engine

import (
	"strconv"

	"github.com/dbriemann/chesskimo/base"
	"github.com/dbriemann/chesskimo/fen"
)

// Board contains all information for a chess board state, including the board itself as 0x88 board.
//
//   +------------------------+
// 8 |70 71 72 73 74 75 76 77 | 78 79 7a 7b 7c 7d 7e 7f
// 7 |60 61 62 63 64 65 66 67 | 68 69 6a 6b 6c 6d 6e 6f
// 6 |50 51 52 53 54 55 56 57 | 58 59 5a 5b 5c 5d 5e 5f
// 5 |40 41 42 43 44 45 46 47 | 48 49 4a 4b 4c 4d 4e 4f
// 4 |30 31 32 33 34 35 36 37 | 38 39 3a 3b 3c 3d 3e 3f
// 3 |20 21 22 23 24 25 26 27 | 28 29 2a 2b 2c 2d 2e 2f
// 2 |10 11 12 13 14 15 16 17 | 18 19 1a 1b 1c 1d 1e 1f
// 1 | 0  1  2  3  4  5  6  7 |  8  9  a  b  c  d  e  f
//   +------------------------+
//     a  b  c  d  e  f  g  h
//
// All indexes shown are in HEX. The left board represents the actual playing board, whereas there right board
// is for detection of illegal moves without heavy conditional usage.
type Board struct {
	Squares     [64 * 2]base.Piece
	CastleShort [2]bool
	CastleLong  [2]bool
	IsCheck     bool
	EpSquare    base.Square
	Player      base.Color
	MoveNumber  uint16
	DrawCounter uint16
	Kings       [2]base.Square
	Queens      [2]base.PieceList
	Rooks       [2]base.PieceList
	Bishops     [2]base.PieceList
	Knights     [2]base.PieceList
	Pawns       [2]base.PieceList
}

const ()

// Lookup0x88 maps the indexes of a 8x8 MinBoard to the 0x88 board indexes.
var Lookup0x88 = [64]base.Square{
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
	b.Kings = [2]base.Square{base.OTB, base.OTB}
	b.Queens = [2]base.PieceList{base.NewPieceList()}
	b.Rooks = [2]base.PieceList{base.NewPieceList()}
	b.Bishops = [2]base.PieceList{base.NewPieceList()}
	b.Knights = [2]base.PieceList{base.NewPieceList()}
	b.Pawns = [2]base.PieceList{base.NewPieceList()}

	b.SetStartingPosition()
	return b
}

// SetStartingPosition sets the starting position of a default chess game via FEN code.
func (b *Board) SetStartingPosition() {
	err := b.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}
}

// SetFEN sets a chess position from a 'FEN' string. This function is expensive
// in terms of computation time. It should only be used in non performance critical
// code such as setting up states etc.
func (b *Board) SetFEN(fenstr string) error {
	mb, err := fen.ParseFEN(fenstr)
	if err != nil {
		panic(err)
	}

	for col := base.BLACK; col <= base.WHITE; col++ {
		b.Queens[col].Clear()
		b.Rooks[col].Clear()
		b.Bishops[col].Clear()
		b.Knights[col].Clear()
		b.Pawns[col].Clear()
	}

	// No error encountered -> set all members of the board instance.
	b.MoveNumber = mb.MoveNum
	b.DrawCounter = mb.HalfMoves

	if mb.EpSquare != base.OTB {
		b.EpSquare = Lookup0x88[mb.EpSquare]
	} else {
		b.EpSquare = base.OTB
	}

	b.CastleShort = mb.CastleShort
	b.CastleLong = mb.CastleLong
	b.Player = mb.Color
	for idx, piece := range mb.Squares {
		sq := Lookup0x88[idx]
		b.Squares[sq] = piece
		switch piece {
		case base.WKING:
			b.Kings[base.WHITE] = sq
		case base.BKING:
			b.Kings[base.BLACK] = sq
		case base.WQUEEN:
			b.Queens[base.WHITE].Add(sq)
		case base.BQUEEN:
			b.Queens[base.BLACK].Add(sq)
		case base.WROOK:
			b.Rooks[base.WHITE].Add(sq)
		case base.BROOK:
			b.Rooks[base.BLACK].Add(sq)
		case base.WBISHOP:
			b.Bishops[base.WHITE].Add(sq)
		case base.BBISHOP:
			b.Bishops[base.BLACK].Add(sq)
		case base.WKNIGHT:
			b.Knights[base.WHITE].Add(sq)
		case base.BKNIGHT:
			b.Knights[base.BLACK].Add(sq)
		case base.WPAWN:
			b.Pawns[base.WHITE].Add(sq)
		case base.BPAWN:
			b.Pawns[base.BLACK].Add(sq)
		}
	}

	return nil
}

// GeneratePawnMoves generates all pseudo legal pawn moves for the given color
// and stores them in the given MoveList.
func (b *Board) GeneratePawnMoves(mlist *base.MoveList, color base.Color) {
	from, to := base.OTB, base.OTB
	tpiece := base.NO_PIECE
	oppColor := color.FlipColor()
	move := base.Move{}
	piece := base.PAWN | color

	// Possible en passent captures are detected backwards
	// so we do not need to add another conditional to the
	// capture loop below.
	if b.EpSquare != base.OTB {
		from = b.EpSquare
		// We use the 'wrong' color to find e.p. captures
		// by searching in the opposite direction.
		// A potential overflow of int8(from) + capdir can generally be ignored
		// because casting it back to uint8 will correct the result.
		to = base.Square(int8(from) + base.PAWN_CAPTURE_DIRS[oppColor][0])
		tpiece = b.Squares[to]
		if to.IsLegal() && tpiece.HasColor(color) {
			move = base.NewMove(to, from, piece, base.PAWN|oppColor, base.NO_PIECE, base.EP_TYPE_CAPTURE, base.NONE)
			mlist.Put(move)
		}

		to = base.Square(int8(from) + base.PAWN_CAPTURE_DIRS[oppColor][1])
		tpiece = b.Squares[to]
		if to.IsLegal() && tpiece.HasColor(color) {
			move = base.NewMove(to, from, piece, base.PAWN|oppColor, base.NO_PIECE, base.EP_TYPE_CAPTURE, base.NONE)
			mlist.Put(move)
		}
	}

	for i := uint8(0); i < b.Pawns[color].Size; i++ {
		// 0. Retrieve 'from' square from piece list.
		from = b.Pawns[color].Pieces[i]

		// a. Captures
		for _, capdir := range base.PAWN_CAPTURE_DIRS[color] {
			to = base.Square(int8(from) + capdir)
			tpiece = b.Squares[to]
			// If the target square is on board and has the opponent's color
			// the capture is possible.
			if to.IsLegal() && tpiece.HasColor(oppColor) {
				// We also have to check if the capture is also a promotion.
				if to.IsPawnPromoting(color) {
					for prom := base.QUEEN; prom >= base.KNIGHT; prom >>= 1 {
						move = base.NewMove(from, to, piece, tpiece, prom|color, base.EP_TYPE_NONE, base.NONE)
						mlist.Put(move)
					}
				} else {
					move = base.NewMove(from, to, piece, tpiece, base.NO_PIECE, base.EP_TYPE_NONE, base.NONE)
					mlist.Put(move)
				}
			}
		}

		// b. Push by one.
		// Target square does never need legality check here.
		to = base.Square(int8(from) + base.PAWN_PUSH_DIRS[color])
		if b.Squares[to].IsEmpty() {
			if to.IsPawnPromoting(color) {
				for prom := base.QUEEN; prom >= base.KNIGHT; prom >>= 1 {
					move = base.NewMove(from, to, piece, base.NO_PIECE, prom|color, base.EP_TYPE_NONE, base.NONE)
					mlist.Put(move)
				}
			} else {
				move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_NONE, base.NONE)
				mlist.Put(move)
			}
		}
		// c. Double push by advancing one more time, if the pawn was at base rank.
		if from.IsPawnBaseRank(color) {
			to = base.Square(int8(to) + base.PAWN_PUSH_DIRS[color])
			if b.Squares[to].IsEmpty() {
				move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_CREATE, base.NONE)
				mlist.Put(move)
			}
		}
	}
}

// GenerateKnightMoves generates all pseudo legal knight moves for the given color
// and stores them in the given MoveList.
func (b *Board) GenerateKnightMoves(mlist *base.MoveList, color base.Color) {
	from, to := base.OTB, base.OTB
	tpiece := base.NO_PIECE
	oppColor := color.FlipColor()
	move := base.Move{}
	piece := base.KNIGHT | color

	// Iterate all knights of 'color'.
	for i := uint8(0); i < b.Knights[color].Size; i++ {
		from = b.Knights[color].Pieces[i]
		// Try all possible directions for a knight.
		for _, dir := range base.KNIGHT_DIRS {
			to = base.Square(int8(from) + dir)
			if to.IsLegal() {
				tpiece = b.Squares[to]
				if tpiece.IsEmpty() {
					// Add a normal move.
					move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_NONE, base.NONE)
					mlist.Put(move)
				} else if tpiece.HasColor(oppColor) {
					// Add a capture move.
					move = base.NewMove(from, to, piece, tpiece, base.NO_PIECE, base.EP_TYPE_NONE, base.NONE)
					mlist.Put(move)
				} // Else the square is occupied by own piece.
			} // Else target is outside of board.
		}
	}
}

func (b *Board) String() string {
	str := "  +-----------------+\n"
	for r := 7; r >= 0; r-- {
		str += strconv.Itoa(r+1) + " | "
		for f := 0; f < 8; f++ {
			idx := 16*r + f
			if base.Piece(idx) == b.EpSquare {
				str += ", "
			} else {
				str += base.PrintMap[b.Squares[idx]] + " "
			}

			if f == 7 {
				str += "|\n"
			}
		}
	}
	str += "  +-----------------+\n"
	str += "    a b c d e f g h\n"

	return str
}
