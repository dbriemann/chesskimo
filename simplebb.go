package main

import (
	"fmt"

	"github.com/zensword/chesskimo/fen"
)

// SimpleBB implements a chess game board state with 'kindergarten' bitboards.
type SimpleBB struct {
	bbPiecesByColoruint64 [2]uint64
	bbPiecesByType        [6]uint64
	kings                 [2]uint8
	enPassentSquare       uint8
	drawCounter           uint8
	moveNumber            uint8 // could this be too small?
	player                Color
	castleShort           [2]bool
	castleLong            [2]bool
	isCheck               bool
}

// NewRootBB creates a new instance of a SimpleBB.
// The created state is the root/starting position of a chess game.
func NewRootBB() SimpleBB {
	bb := SimpleBB{}
	bb.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	return bb
}

// SetFEN sets a chess position encoded as 'FEN' string.
func (bb *SimpleBB) SetFEN(fenstr string) error {
	// extract fields
	fields, err := fen.SplitFields(fenstr)
	if err != nil {
		return err
	}

	// extract piece positions
	pieces, err := fen.ParsePieces(fields[0])
	fmt.Println(pieces)

	// extract active color

	// extract castling rights

	// extract en passent target square

	// extract half moves since last capture or pawn movement

	// extract full move number

	return nil
}
