package chesskimo

import (
	"testing"
	"time"
)

func TestPerft(t *testing.T) {
	// Currently test validates all FEN positions with perft until depth of 5.
	depth := 5
	testset := map[string]uint64{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1":             4865609,   // 1. Start position
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1": 193690690, // 2. Good testposition
		"n1n5/PPPk4/8/8/8/8/4Kppp/5N1N b - - 0 1":                              3605103,   // 3. Many Promotions
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 0": 193690690, // 4. Kiwipete
		"r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1":     15833292,  // 5. Dense
		"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 0":                            674624,    // 6. Endgame
	}

	board := NewBoard()

	set := 0
	for fen, result := range testset {
		err := board.SetFEN(fen)
		if err != nil {
			t.Fatalf(err.Error())
		}

		start := time.Now()
		moves := board.Perft(depth)
		end := time.Now()

		if moves != result {
			t.Fatalf("Perft result for FEN %s (test %d) is %d but should be %d at depth %d.\n", fen, set, moves, result, depth)
		} else {
			t.Logf("Perft result for FEN %s (test %d) is %d moves. Time used: %f\n", fen, set, moves, end.Sub(start).Seconds())
		}

		set++
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
		"[Ke1-f1, Ke1-d1, Ke1-e2, 0-0, 0-0-0]",
		"[ke8-f8, ke8-e7, ke8-d7, 0-0]",
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
		"[Qe1xqd1, Qe1-e2, Qe1-e3, Qe1xpe4, Qe1-f1, Qe1-g1, Qe1-h1, Qd3-c3, Qd3-b3, Qd3-a3, Qd3-d4, Qd3xpd5, Qd3-e3, Qd3xpf3, Qd3-d2, Qd3xqd1, Qa4-a5, Qa4-a6, Qa4-a7, Qa4-a8, Qa4xnb4, Qa4-a3, Qa4-a2, Qa4-a1, Qc5-b5, Qc5-a5, Qc5xrc6, Qc5xpd5, Qc5-c4, Qc5-c3, Qc5-c2, Qc5-c1, Qb7-a7, Qb7-b8, Qb7xnb6, Qc7-c8, Qc7xrd7, Qc7xrc6, Qe1-d2, Qe1-c3, Qe1xnb4, Qe1-f2, Qe1-g3, Qe1-h4, Qd3-c4, Qd3-b5, Qd3-a6, Qd3xpe4, Qd3-c2, Qd3-b1, Qd3-e2, Qd3-f1, Qa4-b5, Qa4xrc6, Qa4-b3, Qa4-c2, Qa4xqd1, Qc5xnb6, Qc5xpd6, Qc5xnb4, Qc5-d4, Qc5-e3, Qc5-f2, Qc5-g1, Qb7-a8, Qb7-c8, Qb7-a6, Qb7xrc6, Qc7-b8, Qc7-d8, Qc7xnb6, Qc7xpd6]",
		"[qd8-c8, qd8-b8, qd8-d7, qd8-d6, qd8-d5, qd8xPd4, qd8-c7]",
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
		"[Rh1-h2, Rh1-h3, Rh1-h4, Rh1-h5, Rh1-h6, Rh1xph7]",
		"[Rb1-a1, Rb1xrb2, Rb1-c1, Rb1-d1, Rb1-e1, Rb1xnf1, Re2-d2, Re2-c2, Re2xrb2, Re2-e3, Re2xbe4, Re2-f2, Re2-g2, Re2-h2, Re2-e1, Rc5-b5, Rc5-a5, Rc5xpc6, Rc5-d5, Rc5-e5, Rc5-f5, Rc5-g5, Rc5-h5, Rc5xpc4]",
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
		"[Be3-d4, Be3xbc5, Be3-f4, Be3-g5, Be3-h6, Bc4-b5, Bc4-a6, Bc4-d5, Bc4xbe6, Bc4-b3]",
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
		"[Nb1-a3, Nb1-c3, Ng1-f3, Ng1-h3]",
		"[nf6-g8, nf6xPe4, nf6xPg4, nf6-d7, nf6-h5]",
		"[Nf3xpe5, Nf3-g5, Nf3-g1, Nf3-d4, Nf3-d2, Nf3-h4]",
	}

	board := NewBoard()
	mlist := MoveList{}

	for i, fen := range fens {
		board.SetFEN(fen)
		mlist.Clear()
		board.GenerateKnightMoves(&mlist, board.Player)
		strmoves := mlist.String()
		if strmoves != results[i] {
			t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, results[i], &mlist)
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
		"[Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pd2-d3, Pd2-d4, Pe2-e3, Pe2-e4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Ph2-h3, Ph2-h4]",
		"[Pe5xpd6 e.p., Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Ph2-h3, Ph2-h4, Pd4xpc5, Pe5xnf6, Pe5-e6]",
		"[Pe5xpf6 e.p., Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Pg7xbf8=Q, Pg7xbf8=R, Pg7xbf8=B, Pg7xbf8=N, Pg7xrh8=Q, Pg7xrh8=R, Pg7xrh8=B, Pg7xrh8=N]",
		"[Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Pg7xbf8=Q, Pg7xbf8=R, Pg7xbf8=B, Pg7xbf8=N, Pg7xrh8=Q, Pg7xrh8=R, Pg7xrh8=B, Pg7xrh8=N, Pg7-g8=Q, Pg7-g8=R, Pg7-g8=B, Pg7-g8=N]",
		"[pc4xPb3 e.p., ph2xNg1=q, ph2xNg1=r, ph2xNg1=b, ph2xNg1=n, ph2-h1=q, ph2-h1=r, ph2-h1=b, ph2-h1=n, pc4-c3, pc6-c5, pa7-a6, pa7-a5, ph7-h6, ph7-h5]",
		"[Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pd2-d3, Pd2-d4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Ph2-h3, Ph2-h4]",
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
