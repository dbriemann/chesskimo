package fen

import (
	"errors"
	"strings"
)

var (
	// ErrFENFieldsNotSix indicates that a FEN record has the wrong amount of fields.
	// A valid FEN must have 6 fields.
	ErrFENFieldsNotSix = errors.New("FEN has wrong amount of fields (must be 6)")
)

// SplitFields splits a FEN into its fields and returns them seperated into a slice,
// or an error if the amount of fields is not equal 6.
func SplitFields(fen string) ([]string, error) {
	fields := strings.Split(fen, " ")
	if len(fields) != 6 {
		return nil, ErrFENFieldsNotSix
	}

	return fields, nil
}

// ParsePieces parses the 'pieces' string, which is the first
// field of a FEN record and returns a simple board as slice.
func ParsePieces(pieces string) ([]int, error) {
	board := []int{}

	return board, nil
}
