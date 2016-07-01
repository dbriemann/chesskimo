package main

// SimpleBB implements a chess game board state with 'kindergarten' bitboards.
type SimpleBB struct {
	bbPiecesByColoruint64 [2]uint64
	bbPiecesByType        [6]uint64
	kings                 [2]uint8
	enPassentSquare       uint8
	drawCounter           uint8
	moveNumber            uint8 // could this be too small?
	player                uint8
	castleShort           [2]bool
	castleLong            [2]bool
	isCheck               bool
}
