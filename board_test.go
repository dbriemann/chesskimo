package chesskimo

import (
	"testing"
	"time"
)

// BenchmarkPerft
// Used to show go1.10 performance loss compared to go1.9.4
//
// Funnily in this benchmark there is no difference opposed
// to /cmd/bench which is ~10% slower with go1.10.
//
// TODO -> What's going on here?
//
func BenchmarkPerft(b *testing.B) {
	type set struct {
		Fen    string
		Depth  int
		Result uint64
	}
	testsets := []set{
		set{Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Depth: 6, Result: 119060324},
		set{Fen: "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", Depth: 5, Result: 193690690},
		set{Fen: "n1n5/PPPk4/8/8/8/8/4Kppp/5N1N b - - 0 1", Depth: 6, Result: 71179139},
		set{Fen: "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1", Depth: 5, Result: 15833292},
		set{Fen: "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 0", Depth: 7, Result: 178633661},
	}

	board := NewBoard()

	for n := 0; n < b.N; n++ {
		for i, set := range testsets {
			board.SetFEN(set.Fen)
			nodes := board.Perft(set.Depth)
			if nodes != set.Result {
				b.Fatalf("Perft set %d failed", i)
			}
		}
	}
}

//
//
//

func TestPerft(t *testing.T) {
	// Currently test validates all FEN positions with perft until depth of 5.
	depth := 5
	testset := map[string]uint64{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1":             4865609,   // 1. Start position
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1": 193690690, // 2. Good testposition
		"n1n5/PPPk4/8/8/8/8/4Kppp/5N1N b - - 0 1":                              3605103,   // 3. Many Promotions
		"r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1":     15833292,  // 5. Dense
		"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 0":                            674624,    // 6. Endgame
	}

	board := NewBoard()

	for fen, result := range testset {
		t.Log("FEN=", fen)
		err := board.SetFEN(fen)
		if err != nil {
			t.Fatalf(err.Error())
		}

		start := time.Now()
		moves := board.Perft(depth)
		end := time.Now()

		if moves != result {
			t.Fatalf("Perft result for FEN %s is %d but should be %d at depth %d.\n", fen, moves, result, depth)
		} else {
			t.Logf("Perft result for FEN %s is %d moves. Time used: %f\n", fen, moves, end.Sub(start).Seconds())
		}
	}
}

func TestSquareDiffs(t *testing.T) {
	type set struct {
		From   Square
		To     Square
		Pieces Piece
	}

	testset := []set{
		set{0x72, 0x12, ROOK | QUEEN},          // file test down to up
		set{0x12, 0x72, ROOK | QUEEN},          // file test up to down
		set{0x51, 0x54, ROOK | QUEEN},          // rank test left to right
		set{0x54, 0x51, ROOK | QUEEN},          // rank test right to left
		set{0x64, 0x37, BISHOP | QUEEN},        // diagonal up_left to down_right test
		set{0x37, 0x64, BISHOP | QUEEN},        // diagonal down_right to up_left test
		set{0x77, 0x0, BISHOP | QUEEN},         // diagonal up_right to down_left test
		set{0x0, 0x77, BISHOP | QUEEN},         // diagonal down_left to up_right test
		set{0x0, 0x12, KNIGHT},                 // knight jump right+up_right
		set{0x0, 0x21, KNIGHT},                 // knight jump up+up_right
		set{0x77, 0x56, KNIGHT},                // knight jump down+down_left
		set{0x77, 0x65, KNIGHT},                // knight jump left+down_left
		set{0x74, 0x65, KING | QUEEN | BISHOP}, // one step diagonal (king test)
		set{0x4, 0x5, KING | QUEEN | ROOK},     // one step orthogonal (king test)
		// impossible moves
		set{0x0, 0x71, NONE},
		set{0x5, 0x36, NONE},
	}

	for i, ts := range testset {
		diff := 0x77 + ts.From - ts.To
		lookup := SQUARE_DIFFS[diff]
		if !(lookup == ts.Pieces) {
			t.Fatalf("Test %d from %s to %s should be reachable by pieces %d but result ist %d\n", i, PrintBoardIndex[ts.From], PrintBoardIndex[ts.To], ts.Pieces, lookup)
		}
	}
}

func TestDiffDirs(t *testing.T) {
	type set struct {
		From Square
		To   Square
		Dir  int8
	}

	testset := []set{
		set{0x70, 0x07, DOWN_RIGHT},
		set{0x47, 0x65, UP_LEFT},
		set{0x77, 0x11, DOWN_LEFT},
		set{0x24, 0x46, UP_RIGHT},
		set{0x61, 0x66, RIGHT},
		set{0x33, 0x30, LEFT},
		set{0x22, 0x62, UP},
		set{0x55, 0x25, DOWN},
		set{0x00, 0x12, RIGHT + UP_RIGHT},
		set{0x00, 0x21, UP + UP_RIGHT},
		set{0x70, 0x62, RIGHT + DOWN_RIGHT},
		set{0x70, 0x51, DOWN + DOWN_RIGHT},
		set{0x77, 0x65, LEFT + DOWN_LEFT},
		set{0x77, 0x56, DOWN + DOWN_LEFT},
		set{0x07, 0x15, LEFT + UP_LEFT},
		set{0x07, 0x26, UP + UP_LEFT},
	}

	for i, ts := range testset {
		diff := 0x77 + ts.From - ts.To
		lookup := DIFF_DIRS[diff]
		if !(lookup == ts.Dir) {
			t.Fatalf("Test %d from %s to %s should be dir %d but result ist %d\n", i, PrintBoardIndex[ts.From], PrintBoardIndex[ts.To], ts.Dir, lookup)
		}
	}
}

func TestGenerateKingMoves(t *testing.T) {
	fens := []string{
		"r3k2r/pppq1ppp/2npbn2/2b1p3/2B1P3/2NPBN2/PPPQ1PPP/R3K2R w KQkq - 4 8",
		"r1bqk2r/ppp2ppp/2n2n2/1B1pp3/1b1PP1P1/2N2N2/PPP2P1P/R1BQK2R b KQkq g3 0 6",
	}
	results := []string{
		"[e1f1, e1d1, e1e2, e1g1, e1c1]",
		"[e8f8, e8e7, e8d7, e8g8]",
	}

	board := NewBoard()
	mlist := MoveList{}

	for i, fen := range fens {
		board.SetFEN(fen)
		mlist.Clear()
		board.GenerateKingMoves(&mlist, board.Player)
		strmoves := mlist.String()
		if strmoves != results[i] {
			t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, results[i], &mlist)
		}
	}
}

func TestGenerateQueenMoves(t *testing.T) {
	fens := []string{
		"7k/1QQr4/1nrp4/2Qp4/Qn2pp2/3Q1p2/1K4p1/3qQ3 w - - 0 1",
		"r2qkbnr/pp2p1Pp/1np1b3/4P3/1PpP1P2/8/P6p/RNBQKBN1 b Qkq b3 0 12",
	}
	results := []string{
		"[e1d1, e1e2, e1e3, e1e4, e1f1, e1g1, e1h1, d3c3, d3b3, d3a3, d3d4, d3d5, d3e3, d3f3, d3d2, d3d1, a4a5, a4a6, a4a7, a4a8, a4b4, a4a3, a4a2, a4a1, c5b5, c5a5, c5c6, c5d5, c5c4, c5c3, c5c2, c5c1, b7a7, b7b8, b7b6, c7c8, c7d7, c7c6, e1d2, e1c3, e1b4, e1f2, e1g3, e1h4, d3c4, d3b5, d3a6, d3e4, d3c2, d3b1, d3e2, d3f1, a4b5, a4c6, a4b3, a4c2, a4d1, c5b6, c5d6, c5b4, c5d4, c5e3, c5f2, c5g1, b7a8, b7c8, b7a6, b7c6, c7b8, c7d8, c7b6, c7d6]",
		"[d8c8, d8b8, d8d7, d8d6, d8d5, d8d4, d8c7]",
	}

	board := NewBoard()
	mlist := MoveList{}

	for i, fen := range fens {
		board.SetFEN(fen)
		mlist.Clear()
		board.GenerateQueenMoves(&mlist, board.Player)
		strmoves := mlist.String()
		if strmoves != results[i] {
			t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, results[i], &mlist)
		}
	}
}

func TestGenerateRookMoves(t *testing.T) {
	fens := []string{
		"r2qkbnr/pp2p1Pp/1np1b3/3pPp2/3P4/8/PPP2PP1/RNBQKBNR w KQkq f6 0 8",
		"8/5k2/2p5/2R5/2p1b3/2K5/1r2R3/1R3n2 w - - 0 1",
	}
	results := []string{
		"[h1h2, h1h3, h1h4, h1h5, h1h6, h1h7]",
		"[b1a1, b1b2, b1c1, b1d1, b1e1, b1f1, e2d2, e2c2, e2b2, e2e3, e2e4, e2f2, e2g2, e2h2, e2e1, c5b5, c5a5, c5c6, c5d5, c5e5, c5f5, c5g5, c5h5, c5c4]",
	}

	board := NewBoard()
	mlist := MoveList{}

	for i, fen := range fens {
		board.SetFEN(fen)
		mlist.Clear()
		board.GenerateRookMoves(&mlist, board.Player)
		strmoves := mlist.String()
		if strmoves != results[i] {
			t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, results[i], &mlist)
		}
	}
}

func TestGenerateBishopMoves(t *testing.T) {
	fens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"r3k2r/pppq1ppp/2npbn2/2b1p3/2B1P3/2NPBN2/PPPQ1PPP/R3K2R w KQkq - 4 8",
	}
	results := []string{
		"[]",
		"[e3d4, e3c5, e3f4, e3g5, e3h6, c4b5, c4a6, c4d5, c4e6, c4b3]",
	}

	board := NewBoard()
	mlist := MoveList{}

	for i, fen := range fens {
		board.SetFEN(fen)
		mlist.Clear()
		board.GenerateBishopMoves(&mlist, board.Player)
		strmoves := mlist.String()
		if strmoves != results[i] {
			t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, results[i], &mlist)
		}
	}
}

func TestGenerateKnightMoves(t *testing.T) {
	fens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"r1bqk2r/ppp2ppp/2n2n2/1B1pp3/1b1PP1P1/2N2N2/PPP2P1P/R1BQK2R b KQkq g3 0 6",
		"r1bqk2r/pppp1pp1/2n2n1p/1B2p3/1b2P3/2NP1N2/PPP2PPP/R1BQK2R w KQkq - 2 6",
	}
	results := []string{
		"[b1a3, b1c3, g1f3, g1h3]",
		"[f6g8, f6e4, f6g4, f6d7, f6h5]",
		"[f3e5, f3g5, f3g1, f3d4, f3d2, f3h4]",
	}

	board := NewBoard()
	mlist := MoveList{}

	for i, fen := range fens {
		board.SetFEN(fen)
		mlist.Clear()
		board.GenerateKnightMoves(&mlist, board.Player)
		strmoves := mlist.String()
		if strmoves != results[i] {
			t.Fatalf("Position\n %s %s expected move list: %s\n but got: %s\n", &board, board.InfoBoardString(), results[i], &mlist)
		}
	}
}

func TestGeneratePawnMoves(t *testing.T) {
	fens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"rnbqkb1r/pp2pppp/5n2/2ppP3/3P4/8/PPP2PPP/RNBQKBNR w KQkq d6 0 4",
		"r2qkbnr/pp2p1Pp/1np1b3/3pPp2/3P4/8/PPP2PP1/RNBQKBNR w KQkq f6 0 8",
		"r2qkb1r/pp2p1Pp/1np1bn2/3p4/3P4/8/PPP2PP1/RNBQKBNR w KQkq - 0 9",
		"r2qkbnr/pp2p1Pp/1np1b3/4P3/1PpP1P2/8/P6p/RNBQKBN1 b Qkq b3 0 12",
		"rnb1kb1r/ppppn1pp/4p3/2K1Ppq1/8/8/PPPP1PPP/RNBQ1BNR w kq f6 0 7",
	}
	results := []string{
		"[a2a3, a2a4, b2b3, b2b4, c2c3, c2c4, d2d3, d2d4, e2e3, e2e4, f2f3, f2f4, g2g3, g2g4, h2h3, h2h4]",
		"[e5d6, a2a3, a2a4, b2b3, b2b4, c2c3, c2c4, f2f3, f2f4, g2g3, g2g4, h2h3, h2h4, d4c5, e5f6, e5e6]",
		"[e5f6, a2a3, a2a4, b2b3, b2b4, c2c3, c2c4, f2f3, f2f4, g2g3, g2g4, g7f8q, g7f8r, g7f8b, g7f8n, g7h8q, g7h8r, g7h8b, g7h8n]",
		"[a2a3, a2a4, b2b3, b2b4, c2c3, c2c4, f2f3, f2f4, g2g3, g2g4, g7f8q, g7f8r, g7f8b, g7f8n, g7h8q, g7h8r, g7h8b, g7h8n, g7g8q, g7g8r, g7g8b, g7g8n]",
		"[c4b3, h2g1q, h2g1r, h2g1b, h2g1n, h2h1q, h2h1r, h2h1b, h2h1n, c4c3, c6c5, a7a6, a7a5, h7h6, h7h5]",
		"[a2a3, a2a4, b2b3, b2b4, c2c3, c2c4, d2d3, d2d4, f2f3, f2f4, g2g3, g2g4, h2h3, h2h4]",
	}

	board := NewBoard()
	mlist := MoveList{}

	for i, fen := range fens {
		board.SetFEN(fen)
		mlist.Clear()
		board.GeneratePawnMoves(&mlist, board.Player)
		strmoves := mlist.String()
		if strmoves != results[i] {
			t.Fatalf("Position\n %s \n test %d\n expected move list: %s\n but got           : %s\n", &board, i, results[i], &mlist)
		}
	}
}
