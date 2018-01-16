package board

import (
	"strconv"

	"github.com/dbriemann/chesskimo/base"
	"github.com/dbriemann/chesskimo/fen"
)

// SimpleBB implements a chess game board state with 'kindergarten' bitboards.
type SimpleBB struct {
	bbPiecesByColor [2]base.BB
	bbPiecesByType  [6]base.BB
	kings           [2]base.Square
	enPassentSquare base.Square
	drawCounter     uint16
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

func (bb *SimpleBB) String() string {
	str := ""
	rank := 8
	str += "\n +--------+"
	for idx := base.Square(0); idx < 64; idx++ {
		if idx%8 == 0 {
			str += "\n" + strconv.Itoa(rank) + "|"
			rank--
		}
		p := bb.getIdx(idx)
		switch p {
		case base.WPAWN:
			str += "P"
		case base.WKNIGHT:
			str += "N"
		case base.WBISHOP:
			str += "B"
		case base.WROOK:
			str += "R"
		case base.WQUEEN:
			str += "Q"
		case base.WKING:
			str += "K"
		case base.BPAWN:
			str += "p"
		case base.BKNIGHT:
			str += "n"
		case base.BBISHOP:
			str += "b"
		case base.BROOK:
			str += "r"
		case base.BQUEEN:
			str += "q"
		case base.BKING:
			str += "k"
		default:
			str += "."
		}
		if idx%8 == 7 {
			str += "|"
		}
	}
	str += "\n +--------+"
	str += "\n  abcdefgh\n"

	return str
}

// SetStartingPosition sets the starting position of a default chess game via FEN code.
func (bb *SimpleBB) SetStartingPosition() {
	err := bb.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}
}

func (bb *SimpleBB) getSquare(file, rank base.Square) base.Piece {
	idx := 8*rank + file
	return bb.getIdx(idx)
}

func (bb *SimpleBB) setSquare(file, rank base.Square, piece base.Piece) {
	idx := 8*rank + file
	bb.setIdx(idx, piece)
}

func (bb *SimpleBB) getIdx(idx base.Square) base.Piece {
	mask := base.IdxBitmask(base.BB(idx))
	black := bb.bbPiecesByColor[base.BLACK] & mask
	colorOffset := 0
	piece := base.EMPTY

	if black != 0 {
		colorOffset = base.MASK_COLOR
	}

	for p := base.PAWN; p <= base.KING; p++ {
		if (bb.bbPiecesByType[p] & mask) != 0 {
			piece = p
			break
		}
	}

	return piece | base.Piece(colorOffset)
}

func (bb *SimpleBB) setIdx(idx base.Square, piece base.Piece) {
	if piece == base.EMPTY {
		// If the square is empty, delete all bits in all corresponding masks.
		rmask := ^base.IdxBitmask(base.BB(idx))
		bb.bbPiecesByColor[base.WHITE] &= rmask
		bb.bbPiecesByColor[base.BLACK] &= rmask
		for p := base.PAWN; p <= base.KING; p++ {
			bb.bbPiecesByType[p] &= rmask
		}
	} else {
		// Else add the specific bit to the corresponding masks.
		amask := base.IdxBitmask(base.BB(idx))
		bb.bbPiecesByType[piece&base.MASK_PIECE] |= amask
		color := (piece & base.MASK_COLOR) >> 3
		bb.bbPiecesByColor[color] |= amask
		if (piece & base.MASK_PIECE) == base.KING {
			bb.kings[color] = idx
		}
	}
}

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

	// Extract active color.
	color, err := fen.ParseColor(fields[1])
	if err != nil {
		return err
	}

	// Extract castling rights.
	short, long := fen.ParseCastlingRights(fields[2])

	// Extract en passent target square.
	epSq, err := fen.ParseEnPassent(fields[3])
	if err != nil {
		return err
	}

	// Extract half moves since last capture or pawn movement.
	halfMoves, err := fen.ParseMoveNumber(fields[4])
	if err != nil {
		return err
	}

	// Extract full move number.
	moveNum, err := fen.ParseMoveNumber(fields[5])
	if err != nil {
		return err
	}

	// No error encountered -> set all members of the board instance.
	bb.moveNumber = moveNum
	bb.drawCounter = halfMoves
	bb.enPassentSquare = epSq
	bb.castleLong = long
	bb.castleShort = short
	bb.player = color
	for idx, piece := range pieces {
		bb.setIdx(base.Square(idx), piece)
	}

	return nil
}
