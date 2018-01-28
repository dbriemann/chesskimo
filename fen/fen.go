package fen

import (
	"errors"
	"strconv"
	"strings"

	"github.com/dbriemann/chesskimo/base"
)

var (
	// ErrFENFieldsInvalid indicates that a FEN record has invalid fields encoding.
	ErrFENFieldsInvalid = errors.New("FEN has invalid fields encoding")
	// ErrFENRanksInvalid indicates that a FEN record has an invalid encoding.
	ErrFENRanksInvalid = errors.New("FEN has invalid rank encoding")
	// ErrFENColorInvalid indicates that a FEN record has an invalid color encoding.
	ErrFENColorInvalid = errors.New("FEN has invalid color field")
	// ErrFENSquareInvalid indicates that a FEN record contains an invalid square.
	ErrFENSquareInvalid = errors.New("FEN has invalid square")
	// ErrFENMoveNumInvalid indicates that a FEN record contains an invalid move number.
	ErrFENMoveNumInvalid = errors.New("FEN has invalid move number")

	// FENMap maps FEN piece symbols (plus empty ' ') to the internal definition.
	FENMap = map[rune]base.Piece{
		'P': base.WPAWN,
		'N': base.WKNIGHT,
		'B': base.WBISHOP,
		'R': base.WROOK,
		'Q': base.WQUEEN,
		'K': base.WKING,
		'p': base.BPAWN,
		'n': base.BKNIGHT,
		'b': base.BBISHOP,
		'r': base.BROOK,
		'q': base.BQUEEN,
		'k': base.BKING,
		' ': base.EMPTY,
	}
)

// // ValidateFEN validates a FEN string.
// TODO
// func ValidateFEN(fen string) bool {
// 	// Must have 6 fields, split by ONE space.
// 	fields := strings.Split(fen, " ")
// 	if len(fields) != 6 {
// 		return ErrFENFieldsInvalid
// 	}

// 	// Field 0 contains the pieces setup of the board.
// 	// Must have 8 ranks and only specific characters are allowed inside.
// 	rank := "[prnbqkPRNBQK12345678]+"
// 	pattern := fmt.Sprintf("(%s)/(%s)/(%s)/(%s)/(%s)/(%s)/(%s)/(%s)", rank, rank, rank, rank, rank, rank, rank, rank)
// 	reg := regexp.MustCompile(pattern)

// 	return true
// }

func ParseFEN(fen string) (base.MinBoard, error) {
	mb := base.MinBoard{}

	// Extract fields.
	fields, err := splitFENFields(fen)
	if err != nil {
		return mb, err
	}

	// Extract piece positions.
	pieces, err := parseFENPieces(fields[0])
	if err != nil {
		return mb, err
	}

	// Extract active color.
	color, err := parseFENColor(fields[1])
	if err != nil {
		return mb, err
	}

	// Extract castling rights.
	short, long := parseFENCastlingRights(fields[2])

	// Extract en passent target square.
	epSq, err := parseFENEnPassent(fields[3])
	if err != nil {
		return mb, err
	}

	// Extract half moves since last capture or pawn movement.
	halfMoves, err := parseFENMoveNumber(fields[4])
	if err != nil {
		return mb, err
	}

	// Extract full move number.
	moveNum, err := parseFENMoveNumber(fields[5])
	if err != nil {
		return mb, err
	}

	mb.Squares = pieces
	mb.Color = color
	mb.CastleShort = short
	mb.CastleLong = long
	mb.EpSquare = epSq
	mb.HalfMoves = halfMoves
	mb.MoveNum = moveNum

	return mb, nil
}

// SplitFields splits a FEN into its fields and returns them separated into a slice,
// or an error if the amount of fields is not equal 6.
func splitFENFields(fen string) ([]string, error) {
	// Split for any number of whitespaces. Is fault tolerant to some malformed FENs.
	fields := strings.Fields(fen)
	if len(fields) != 6 {
		return nil, ErrFENFieldsInvalid
	}

	return fields, nil
}

// ParsePieces parses the 'pieces' string, which is the first
// field of a FEN record and returns a simple board as slice.
// The returned board is a 1D slice and has the following structure:
//
//    a b c d e f g h
//  +----------------
// 8|56,..         63] <- array index 63 is black rook @ H8
// 7|48,..
// 6|40,..
// 5|32,..
// 4|24,..
// 3|16,..
// 2| 8,..
// 1|[0,1,2,3,4,5,6,7, <- array index 0 is white rook @ A1
//  +-----------------
//
func parseFENPieces(pieces string) ([64]base.Piece, error) {
	board := [64]base.Piece{}

	// Reverse ranks to transform from FEN to internal representation.
	ranks := strings.Split(pieces, "/")
	if len(ranks) != 8 {
		return board, ErrFENRanksInvalid
	}
	ranks[0], ranks[7] = ranks[7], ranks[0]
	ranks[1], ranks[6] = ranks[6], ranks[1]
	ranks[2], ranks[5] = ranks[5], ranks[2]
	ranks[3], ranks[4] = ranks[4], ranks[3]

	// Put it back together so old code below can just continue to work.
	pieces = ranks[0] + ranks[1] + ranks[2] + ranks[3] + ranks[4] + ranks[5] + ranks[6] + ranks[7]

	// Replace numbers with an equal amount of spaces
	// so we have 64 characters in total, mapping to
	// a whole board of chess.
	for i := 0; i <= 8; i++ {
		// replace all possibilities of numbers (1-8)
		pieces = strings.Replace(pieces, strconv.Itoa(i), strings.Repeat(" ", i), -1)
	}

	// Check if resulting board is valid in its size.
	if len(pieces) != 64 {
		return board, ErrFENRanksInvalid
	}

	for i, r := range pieces {
		// Test if the FEN code contains a valid piece symbol at position 'i'.
		if piece, ok := FENMap[rune(r)]; ok {
			board[i] = piece
		} else {
			// Unknown symbol..
			return board, ErrFENRanksInvalid
		}
	}

	return board, nil
}

// ParseColor parses the active color from the 'color' string.
func parseFENColor(color string) (base.Color, error) {
	// color = strings.ToLower(color)
	switch color {
	case "w":
		return base.WHITE, nil
	case "b":
		return base.BLACK, nil
	default:
		return 0, ErrFENColorInvalid
	}
}

// ParseCastlingRights parses which player still has rights to castle short and long.
// The returned arrays describe the castling rights as follows:
// short:[BLACK, WHITE], long:[BLACK, WHITE]
func parseFENCastlingRights(castle string) ([2]bool, [2]bool) {
	short := [2]bool{strings.Contains(castle, "k"), strings.Contains(castle, "K")}
	long := [2]bool{strings.Contains(castle, "q"), strings.Contains(castle, "Q")}

	// Invalid characters are ignored and treated as if no player had castling rights.

	return short, long
}

// ParseEnPassent parses the current possible en passent capture square, if there is one.
func parseFENEnPassent(ep string) (base.Square, error) {
	if ep == "-" {
		return base.NONE, nil
	}
	sq, err := parseFENSquare(ep)
	return sq, err
}

// ParseMoveNumber parses a move number from the 'num' string.
func parseFENMoveNumber(num string) (uint16, error) {
	n, err := strconv.ParseInt(num, 10, 16)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, ErrFENMoveNumInvalid
	}

	return uint16(n), nil
}

func parseFENSquare(sq string) (base.Square, error) {
	if len(sq) != 2 {
		return base.NONE, ErrFENSquareInvalid
	}
	r, f := sq[0], sq[1]
	rank, file := r-97, 8-(f-48)

	idx := base.Square(rank + file*8)
	if idx < 0 || idx > 63 {
		return base.NONE, ErrFENSquareInvalid
	}

	return idx, nil
}
