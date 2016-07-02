// Package fen provides basic functionality to decode and encode FEN strings.
package fen

import (
	"errors"
	"strconv"
	"strings"

	"github.com/zensword/chesskimo/base"
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

// SplitFields splits a FEN into its fields and returns them separated into a slice,
// or an error if the amount of fields is not equal 6.
func SplitFields(fen string) ([]string, error) {
	fields := strings.Split(fen, " ")
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
// 8|[0,1,2,3,4,5,6,7, <- array index 0 is black rook @ A8
// 7| 8,..
// 6|16,..
// 5|24,..
// 4|32,..
// 3|40,..
// 2|48,..
// 1|56,..         64] <- array index 63 is white rook @ H1
//  +-----------------
//
func ParsePieces(pieces string) ([64]base.Piece, error) {
	board := [64]base.Piece{}

	// Remove all '/' from the encoded string.
	pieces = strings.Replace(pieces, "/", "", 8)

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
func ParseColor(color string) (base.Color, error) {
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
// short:[WHITE, BLACK], long:[WHITE, BLACK]
func ParseCastlingRights(castle string) ([2]bool, [2]bool) {
	short := [2]bool{strings.Contains(castle, "K"), strings.Contains(castle, "k")}
	long := [2]bool{strings.Contains(castle, "Q"), strings.Contains(castle, "q")}

	// Invalid characters are ignored and treated as if no player had castling rights.

	return short, long
}

// ParseEnPassent parses the current possible en passent capture square, if there is any.
func ParseEnPassent(ep string) (int, error) {
	if ep == "-" {
		return base.NONE, nil
	}
	sq, err := parseSquare(ep)
	return sq, err
}

// ParseMoveNumber parses a move number from the 'num' string.
func ParseMoveNumber(num string) (uint8, error) {
	n, err := strconv.ParseInt(num, 10, 8)
	if err != nil || n < 0 {
		return 0, ErrFENMoveNumInvalid
	}

	return uint8(n), nil
}

func parseSquare(sq string) (int, error) {
	if len(sq) != 2 {
		return base.NONE, ErrFENSquareInvalid
	}
	r, f := sq[0], sq[1]
	rank, file := r-97, 8-(f-48)

	idx := int(rank + file*8)
	if idx < 0 || idx > 63 {
		return base.NONE, ErrFENSquareInvalid
	}

	return idx, nil
}
