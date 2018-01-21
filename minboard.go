package chesskimo

type MinBoard struct {
	Squares     [64]Piece
	Color       Color
	CastleShort [2]bool
	CastleLong  [2]bool
	EpSquare    Square
	HalfMoves   uint16
	MoveNum     uint16
}
