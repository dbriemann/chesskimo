package chesskimo

type Board struct {
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

	//	// No error encountered -> set all members of the board instance.
	//	bb.moveNumber = moveNum
	//	bb.drawCounter = halfMoves
	//	bb.enPassentSquare = epSq
	//	bb.castleLong = long
	//	bb.castleShort = short
	//	bb.player = color
	//	for idx, piece := range pieces {
	//		bb.setIdx(base.Square(idx), piece)
	//	}

	return nil
}

func (b *Board) String() string {
	str := ""
	return str
}
