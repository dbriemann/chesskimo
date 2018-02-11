package chesskimo

import (
	"testing"
)

// TestSplitOK tests if the SplitFields function behaves correctly.
func TestSplitFields(t *testing.T) {
	sixFields := "1 2 3 4 5 6"
	lessThanSixFields := "1   2  3  4"
	moreThanSixFields := "1 2 3 4 5 6 7 8"

	_, err1 := splitFENFields(sixFields)
	if err1 != nil {
		t.Fatalf("Expected pass, split FEN is %s\n", sixFields)
	}
	_, err2 := splitFENFields(lessThanSixFields)
	if err2 == nil {
		t.Fatalf("Expected fail, split FEN is %s\n", lessThanSixFields)
	}

	_, err3 := splitFENFields(moreThanSixFields)
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
		_, err := parseFENPieces(s)
		if err != nil {
			t.Fatalf("Expected pass, pieces FEN is %s\n", s)
		}
	}

	// Should fail.
	for _, s := range invalid {
		_, err := parseFENPieces(s)
		if err == nil {
			t.Fatalf("Expected fail, pieces FEN is %s\n", s)
		}
	}
}

// TestParseColor tests if the ParseColor function behaves correctly.
func TestParseColor(t *testing.T) {
	valid := map[string]Color{
		"w": WHITE,
		"b": BLACK,
	}
	invalid := []string{" ", "-", " w", " b", "w ", "b ", "W", "B", "#", "?", "a"}

	// Should succeed.
	for s, col := range valid {
		c, err := parseFENColor(s)
		if err != nil || c != col {
			t.Fatalf("Expected pass, color FEN is: %s\n", s)
		}
	}

	// Should fail.
	for _, s := range invalid {
		_, err := parseFENColor(s)
		if err == nil {
			t.Fatalf("Expected fail, color FEN i: %s\n", s)
		}
	}
}

// TestParseCastlingRights tests if the ParseCastlingRights function behaves correctly.
func TestParseCastlingRights(t *testing.T) {
	tests := map[string][][2]bool{
		"KkQq": [][2]bool{[2]bool{true, true}, [2]bool{true, true}},
		"KQq":  [][2]bool{[2]bool{false, true}, [2]bool{true, true}},
		"kQq":  [][2]bool{[2]bool{true, false}, [2]bool{true, true}},
		"Kkq":  [][2]bool{[2]bool{true, true}, [2]bool{true, false}},
		"KkQ":  [][2]bool{[2]bool{true, true}, [2]bool{false, true}},
		"Kk":   [][2]bool{[2]bool{true, true}, [2]bool{false, false}},
		"Qq":   [][2]bool{[2]bool{false, false}, [2]bool{true, true}},
		"-":    [][2]bool{[2]bool{false, false}, [2]bool{false, false}},
	}

	for s, r := range tests {
		short, long := parseFENCastlingRights(s)
		if r[0] != short || r[1] != long {
			t.Fatalf("Expected %v, %v, but got %v -> castling FEN is %s\n", short, long, r, s)
		}
	}
}

// TestParseEnPassent tests if the ParseEnPassent function behaves correctly.
// Also tests parseSquare.
func TestParseEnPassent(t *testing.T) {
	// No en passent square.
	sq, err := parseFENEnPassent("-")
	if err != nil || sq != OTB {
		t.Fatalf("Expected pass, en passent FEN is: %s\n", "-")
	}
	// Alibi call.
	_, err = parseFENEnPassent(" ")
	if err == nil {
		t.Fatalf("Expected fail, en passent FEN is: %s\n", "-")
	}

	// Test parseSquare success for all squares.
	squares := map[string]Square{
		"a8": 56, "b8": 57, "c8": 58, "d8": 59, "e8": 60, "f8": 61, "g8": 62, "h8": 63,
		"a7": 48, "b7": 49, "c7": 50, "d7": 51, "e7": 52, "f7": 53, "g7": 54, "h7": 55,
		"a6": 40, "b6": 41, "c6": 42, "d6": 43, "e6": 44, "f6": 45, "g6": 46, "h6": 47,
		"a5": 32, "b5": 33, "c5": 34, "d5": 35, "e5": 36, "f5": 37, "g5": 38, "h5": 39,
		"a4": 24, "b4": 25, "c4": 26, "d4": 27, "e4": 28, "f4": 29, "g4": 30, "h4": 31,
		"a3": 16, "b3": 17, "c3": 18, "d3": 19, "e3": 20, "f3": 21, "g3": 22, "h3": 23,
		"a2": 8, "b2": 9, "c2": 10, "d2": 11, "e2": 12, "f2": 13, "g2": 14, "h2": 15,
		"a1": 0, "b1": 1, "c1": 2, "d1": 3, "e1": 4, "f1": 5, "g1": 6, "h1": 7,
	}

	for s, i := range squares {
		sq, err = parseFENSquare(s)
		if err != nil || sq != i {
			t.Fatalf("Expected pass for square FEN: %s\n", s)
		}
	}

	// Test parseSquare fail.
	// TODO: e.g. "z7"
	nonSquares := []string{"h0", "aa3", "", "a", "k55"}
	for _, s := range nonSquares {
		sq, err = parseFENSquare(s)
		if err == nil {
			t.Fatalf("Expected fail for square FEN: %s but got SQ: %d\n", s, sq)
		}
	}
}

// TestParseMoveNumber tests if the ParseMoveNumber function behaves correctly.
func TestParseMoveNumber(t *testing.T) {
	fails := []string{"-1", "44231"}
	for _, n := range fails {
		num, err := parseFENMoveNumber(n)
		if err == nil {
			t.Fatalf("Expected fail for move number FEN: %s but got NUM: %d\n", n, num)
		}
	}

	succeeds := []string{"0", "11", "25001"}
	for _, n := range succeeds {
		_, err := parseFENMoveNumber(n)
		if err != nil {
			t.Fatalf("Expected pass for move number FEN: %s but got ERR: %s\n", n, err.Error())
		}
	}
}
