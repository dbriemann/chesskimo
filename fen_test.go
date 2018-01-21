package chesskimo

import (
	"testing"
)

// TestSplitOK tests if the SplitFields function behaves correctly.
func TestSplitFields(t *testing.T) {
	sixFields := "1 2 3 4 5 6"
	lessThanSixFields := "1   2  3  4"
	moreThanSixFields := "1 2 3 4 5 6 7 8"

	_, err1 := SplitFields(sixFields)
	if err1 != nil {
		t.Fatalf("Expected pass, split FEN is %s\n", sixFields)
	}
	_, err2 := SplitFields(lessThanSixFields)
	if err2 == nil {
		t.Fatalf("Expected fail, split FEN is %s\n", lessThanSixFields)
	}

	_, err3 := SplitFields(moreThanSixFields)
	if err3 == nil {
		t.Fatalf("Expected fail, split FEN is %s\n", moreThanSixFields)
	}
}

// TestParsePieces tests if the ParsePieces function behaves correctly.
func TestParsePieces(t *testing.T) {
	valid := []string{
		"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R",
	}
	invalid := []string{
		// extra rank
		"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR/pppPPppP",
		// Wrong number for empty fields.
		"rnbqkbnr/pp1ppppp/7/2p5/4P3/8/PPPP1PPP/RNBQKBNR",
		// Unknown piece symbol.
		"rnbqkbnr/pp1pppop/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R",
	}

	// Should succeed.
	for _, s := range valid {
		_, err := ParsePieces(s)
		if err != nil {
			t.Fatalf("Expected pass, pieces FEN is %s\n", s)
		}
	}

	// Should fail.
	for _, s := range invalid {
		_, err := ParsePieces(s)
		if err == nil {
			t.Fatalf("Expected fail, pieces FEN is %s\n", s)
		}
	}
}

// TestParseColor tests if the ParseColor function behaves correctly.
func TestParseColor(t *testing.T) {
	valid := map[string]base.Color{
		"w": base.WHITE,
		"b": base.BLACK,
	}
	invalid := []string{" ", "-", " w", " b", "w ", "b ", "W", "B", "#", "?", "a"}

	// Should succeed.
	for s, col := range valid {
		c, err := ParseColor(s)
		if err != nil || c != col {
			t.Fatalf("Expected pass, color FEN is: %s\n", s)
		}
	}

	// Should fail.
	for _, s := range invalid {
		_, err := ParseColor(s)
		if err == nil {
			t.Fatalf("Expected fail, color FEN i: %s\n", s)
		}
	}
}

// TestParseCastlingRights tests if the ParseCastlingRights function behaves correctly.
func TestParseCastlingRights(t *testing.T) {
	tests := map[string][][2]bool{
		"KkQq": [][2]bool{[2]bool{true, true}, [2]bool{true, true}},
		"KQq":  [][2]bool{[2]bool{true, false}, [2]bool{true, true}},
		"kQq":  [][2]bool{[2]bool{false, true}, [2]bool{true, true}},
		"Kkq":  [][2]bool{[2]bool{true, true}, [2]bool{false, true}},
		"KkQ":  [][2]bool{[2]bool{true, true}, [2]bool{true, false}},
		"Kk":   [][2]bool{[2]bool{true, true}, [2]bool{false, false}},
		"Qq":   [][2]bool{[2]bool{false, false}, [2]bool{true, true}},
		"-":    [][2]bool{[2]bool{false, false}, [2]bool{false, false}},
	}

	for s, r := range tests {
		short, long := ParseCastlingRights(s)
		if r[0] != short || r[1] != long {
			t.Fatalf("Expected %v, %v, but got %v -> castling FEN is %s\n", short, long, r, s)
		}
	}
}

// TestParseEnPassent tests if the ParseEnPassent function behaves correctly.
// Also tests parseSquare.
func TestParseEnPassent(t *testing.T) {
	// No en passent square.
	sq, err := ParseEnPassent("-")
	if err != nil || sq != base.NONE {
		t.Fatalf("Expected pass, en passent FEN is: %s\n", "-")
	}
	// Alibi call.
	_, err = ParseEnPassent(" ")
	if err == nil {
		t.Fatalf("Expected fail, en passent FEN is: %s\n", "-")
	}

	// Test parseSquare success for all squares.
	squares := map[string]int{
		"a8": 0, "b8": 1, "c8": 2, "d8": 3, "e8": 4, "f8": 5, "g8": 6, "h8": 7,
		"a7": 8, "b7": 9, "c7": 10, "d7": 11, "e7": 12, "f7": 13, "g7": 14, "h7": 15,
		"a6": 16, "b6": 17, "c6": 18, "d6": 19, "e6": 20, "f6": 21, "g6": 22, "h6": 23,
		"a5": 24, "b5": 25, "c5": 26, "d5": 27, "e5": 28, "f5": 29, "g5": 30, "h5": 31,
		"a4": 32, "b4": 33, "c4": 34, "d4": 35, "e4": 36, "f4": 37, "g4": 38, "h4": 39,
		"a3": 40, "b3": 41, "c3": 42, "d3": 43, "e3": 44, "f3": 45, "g3": 46, "h3": 47,
		"a2": 48, "b2": 49, "c2": 50, "d2": 51, "e2": 52, "f2": 53, "g2": 54, "h2": 55,
		"a1": 56, "b1": 57, "c1": 58, "d1": 59, "e1": 60, "f1": 61, "g1": 62, "h1": 63,
	}

	for s, i := range squares {
		sq, err = parseSquare(s)
		if err != nil || sq != i {
			t.Fatalf("Expected pass for square FEN: %s\n", s)
		}
	}

	// Test parseSquare fail.
	// TODO: e.g. "z7"
	nonSquares := []string{"h0", "aa3", "", "a", "k55"}
	for _, s := range nonSquares {
		sq, err = parseSquare(s)
		if err == nil {
			t.Fatalf("Expected fail for square FEN: %s but got SQ: %d\n", s, sq)
		}
	}
}

// TestParseMoveNumber tests if the ParseMoveNumber function behaves correctly.
func TestParseMoveNumber(t *testing.T) {
	fails := []string{"-1", "44231"}
	for _, n := range fails {
		num, err := ParseMoveNumber(n)
		if err == nil {
			t.Fatalf("Expected fail for move number FEN: %s but got NUM: %d\n", n, num)
		}
	}

	succeeds := []string{"0", "11", "25001"}
	for _, n := range succeeds {
		_, err := ParseMoveNumber(n)
		if err != nil {
			t.Fatalf("Expected pass for move number FEN: %s but got ERR: %s\n", n, err.Error())
		}
	}
}
