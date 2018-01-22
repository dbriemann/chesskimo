package chesskimo

// MinBoard defines a minimalist type of board structure,
// which is used to transport different types of board encodings
// from one to another.
type MinBoard struct {
	Squares     [64]Piece
	Color       Color
	CastleShort [2]bool
	CastleLong  [2]bool
	EpSquare    Square
	HalfMoves   uint16
	MoveNum     uint16
}
