package engine

import (
	"testing"

	"github.com/dbriemann/chesskimo/base"
)

func TestGenerateKnightMoves(t *testing.T) {
	fens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"r1bqk2r/ppp2ppp/2n2n2/1B1pp3/1b1PP1P1/2N2N2/PPP2P1P/R1BQK2R b KQkq g3 0 6",
	}
	results := []string{
		"Nb1-a3, Nb1-c3, Ng1-f3, Ng1-h3",
		"nc6-b8, nc6xPd4, nc6-a5, nc6-e7, nf6-g8, nf6xPe4, nf6xPg4, nf6-d7, nf6-h5",
	}

	board := NewBoard()
	mlist := base.MoveList{}

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
	}
	results := []string{
		"Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pd2-d3, Pd2-d4, Pe2-e3, Pe2-e4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Ph2-h3, Ph2-h4",
		"Pe5xpd6 e.p., Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Ph2-h3, Ph2-h4, Pd4xpc5, Pe5xnf6, Pe5-e6",
		"Pe5xpf6 e.p., Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Pg7xbf8=Q, Pg7xbf8=R, Pg7xbf8=B, Pg7xbf8=N, Pg7xrh8=Q, Pg7xrh8=R, Pg7xrh8=B, Pg7xrh8=N",
		"Pa2-a3, Pa2-a4, Pb2-b3, Pb2-b4, Pc2-c3, Pc2-c4, Pf2-f3, Pf2-f4, Pg2-g3, Pg2-g4, Pg7xbf8=Q, Pg7xbf8=R, Pg7xbf8=B, Pg7xbf8=N, Pg7xrh8=Q, Pg7xrh8=R, Pg7xrh8=B, Pg7xrh8=N, Pg7-g8=Q, Pg7-g8=R, Pg7-g8=B, Pg7-g8=N",
		"pc4xPb3 e.p., ph2xNg1=q, ph2xNg1=r, ph2xNg1=b, ph2xNg1=n, ph2-h1=q, ph2-h1=r, ph2-h1=b, ph2-h1=n, pc4-c3, pc6-c5, pa7-a6, pa7-a5, pb7-b5, ph7-h6, ph7-h5",
	}

	board := NewBoard()
	mlist := base.MoveList{}

	for i, fen := range fens {
		board.SetFEN(fen)
		mlist.Clear()
		board.GeneratePawnMoves(&mlist, board.Player)
		strmoves := mlist.String()
		if strmoves != results[i] {
			t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, results[i], &mlist)
		}
	}
}
