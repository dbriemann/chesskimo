package main

import (
	"fmt"

	"github.com/dbriemann/chesskimo/base"
	"github.com/dbriemann/chesskimo/engine"
)

func main() {
	fens := []string{
		"rnbqkb1r/pp2pppp/5n2/2ppP3/3P4/8/PPP2PPP/RNBQKBNR w KQkq d6 0 4",
		"r2qkbnr/pp2p1Pp/1np1b3/3pPp2/3P4/8/PPP2PP1/RNBQKBNR w KQkq f6 0 8",
		"r2qkb1r/pp2p1Pp/1np1bn2/3p4/3P4/8/PPP2PP1/RNBQKBNR w KQkq - 0 9",
		"r2qkbnr/pp2p1Pp/1np1b3/4P3/1PpP1P2/8/P6p/RNBQKBN1 b Qkq b3 0 12",
		"r1bqk2r/ppp2ppp/2n2n2/1B1pp3/1b1PP1P1/2N2N2/PPP2P1P/R1BQK2R b KQkq g3 0 6",
	}

	b := engine.NewBoard()
	mlist := base.MoveList{}

	for _, fen := range fens {
		b.SetFEN(fen)
		fmt.Println(&b)
		mlist.Clean()
		//		b.GeneratePawnMoves(&mlist, b.Player)
		b.GenerateKnightMoves(&mlist, b.Player)

		fmt.Println("Moves:")
		fmt.Println(&mlist)
	}

}
