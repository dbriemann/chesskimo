package board

import (
	"fmt"

	"github.com/zensword/chesskimo/base"
	"github.com/zensword/chesskimo/fen"
)

// SimpleBB implements a chess game board state with 'kindergarten' bitboards.
type SimpleBB struct {
	bbPiecesByColor [2]uint64
	bbPiecesByType  [6]uint64
	kings           [2]uint8
	enPassentSquare uint8
	drawCounter     uint8
	moveNumber      uint16
	player          base.Color
	castleShort     [2]bool
	castleLong      [2]bool
	isCheck         bool
}

// NewRootBB creates a new instance of a SimpleBB.
// The created state is the root/starting position of a chess game.
func NewRootBB() SimpleBB {
	bb := SimpleBB{}
	bb.SetStartingPosition()

	return bb
}

// SetStartingPosition sets the starting position of a default chess game via FEN code.
func (bb *SimpleBB) SetStartingPosition() {
	err := bb.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}
}

// func (bb *SimpleBB) GetSquare(file, rank base.Square) {

// }

// SetFEN sets a chess position from a 'FEN' string.
func (bb *SimpleBB) SetFEN(fenstr string) error {
	// Extract fields.
	fields, err := fen.SplitFields(fenstr)
	if err != nil {
		return err
	}

	// Extract piece positions.
	pieces, err := fen.ParsePieces(fields[0])
	if err != nil {
		return err
	}
	fmt.Println("PIECES", pieces)

	// Extract active color.
	color, err := fen.ParseColor(fields[1])
	if err != nil {
		return err
	}
	fmt.Println("COLOR", color)

	// Extract castling rights.
	short, long := fen.ParseCastlingRights(fields[2])
	fmt.Println("CASTLE", short, long)

	// Extract en passent target square.
	epSq, err := fen.ParseEnPassent(fields[3])
	if err != nil {
		return err
	}
	fmt.Println("EP", epSq)

	// Extract half moves since last capture or pawn movement.
	halfMove, err := fen.ParseMoveNumber(fields[4])
	if err != nil {
		return err
	}
	fmt.Println("HALF-MOVE", halfMove)

	// Extract full move number.
	moveNum, err := fen.ParseMoveNumber(fields[5])
	if err != nil {
		return err
	}
	fmt.Println("MOVE", moveNum)

	return nil
}
