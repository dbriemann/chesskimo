package engine

import (
	"strings"
	"testing"

	"github.com/dbriemann/chesskimo/base"
)

func TestGeneratePawnMoves(t *testing.T) {
	result1 := "a2-a3, a2-a4, b2-b3, b2-b4, c2-c3, c2-c4, d2-d3, d2-d4, e2-e3, e2-e4, f2-f3, f2-f4, g2-g3, g2-g4, h2-h3, h2-h4"
	result2 := "e5xd6 e.p., a2-a3, a2-a4, b2-b3, b2-b4, c2-c3, c2-c4, f2-f3, f2-f4, g2-g3, g2-g4, h2-h3, h2-h4, d4xc5, e5xnf6, e5-e6"
	result3 := "e5xf6 e.p., a2-a3, a2-a4, b2-b3, b2-b4, c2-c3, c2-c4, f2-f3, f2-f4, g2-g3, g2-g4, g7xbf8=Q, g7xbf8=R, g7xbf8=B, g7xbf8=N, g7xrh8=Q, g7xrh8=R, g7xrh8=B, g7xrh8=N"
	result4 := "a2-a3, a2-a4, b2-b3, b2-b4, c2-c3, c2-c4, f2-f3, f2-f4, g2-g3, g2-g4, g7xbf8=Q, g7xbf8=R, g7xbf8=B, g7xbf8=N, g7xrh8=Q, g7xrh8=R, g7xrh8=B, g7xrh8=N, g7-g8=Q, g7-g8=R, g7-g8=B, g7-g8=N"
	result5 := "c4xb3 e.p., h2xNg1=q, h2xNg1=r, h2xNg1=b, h2xNg1=n, h2-h1=q, h2-h1=r, h2-h1=b, h2-h1=n, c4-c3, c6-c5, a7-a6, a7-a5, b7-b5, h7-h6, h7-h5"

	board := NewBoard()
	mlist := base.MoveList{}

	// Generate pawn moves for root position.
	// There are only single and double pawn pushes here.
	board.GeneratePawnMoves(&mlist, board.Player)
	// Transform move list to a string.
	strmoves := mlist.String()
	// Test resulting moves.
	if strings.Compare(strmoves, result1) != 0 {
		t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, result1, &mlist)
	}

	// The following position includes en passent and other captures.
	mlist.Clean()
	board.SetFEN("rnbqkb1r/pp2pppp/5n2/2ppP3/3P4/8/PPP2PPP/RNBQKBNR w KQkq d6 0 4")
	board.GeneratePawnMoves(&mlist, board.Player)
	strmoves = mlist.String()
	if strings.Compare(strmoves, result2) != 0 {
		t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, result2, &mlist)
	}

	// The following position contains en passent and captures to promotion.
	mlist.Clean()
	board.SetFEN("r2qkbnr/pp2p1Pp/1np1b3/3pPp2/3P4/8/PPP2PP1/RNBQKBNR w KQkq f6 0 8")
	board.GeneratePawnMoves(&mlist, board.Player)
	strmoves = mlist.String()
	if strings.Compare(strmoves, result3) != 0 {
		t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, result3, &mlist)
	}

	// The following position contains captures to promotion and a push to promotion.
	mlist.Clean()
	board.SetFEN("r2qkb1r/pp2p1Pp/1np1bn2/3p4/3P4/8/PPP2PP1/RNBQKBNR w KQkq - 0 9")
	board.GeneratePawnMoves(&mlist, board.Player)
	strmoves = mlist.String()
	if strings.Compare(strmoves, result4) != 0 {
		t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, result4, &mlist)
	}

	// The following position contains captures to promotion and a push to promotion and en passent.
	// This time for the black side.
	mlist.Clean()
	board.SetFEN("r2qkbnr/pp2p1Pp/1np1b3/4P3/1PpP1P2/8/P6p/RNBQKBN1 b Qkq b3 0 12")
	board.GeneratePawnMoves(&mlist, board.Player)
	strmoves = mlist.String()
	if strings.Compare(strmoves, result5) != 0 {
		t.Fatalf("Position\n %s expected move list: %s\n but got: %s\n", &board, result5, &mlist)
	}
}
