package main

import (
	"fmt"

	"github.com/dbriemann/chesskimo/base"
	"github.com/dbriemann/chesskimo/engine"
)

func main() {
	b := engine.NewBoard()
	mlist := base.MoveList{}

	fmt.Println(&b)

	b.GeneratePawnMoves(&mlist, b.Player)

	fmt.Println("Moves:")
	fmt.Println(&mlist)

	b.SetFEN("rnbqkb1r/pp2pppp/5n2/2ppP3/3P4/8/PPP2PPP/RNBQKBNR w KQkq d6 0 4")
	fmt.Println(&b)
	mlist.Clean()
	b.GeneratePawnMoves(&mlist, b.Player)

	fmt.Println("Moves:")
	fmt.Println(&mlist)

	b.SetFEN("r2qkbnr/pp2p1Pp/1np1b3/3pPp2/3P4/8/PPP2PP1/RNBQKBNR w KQkq f6 0 8")
	fmt.Println(&b)
	mlist.Clean()
	b.GeneratePawnMoves(&mlist, b.Player)

	fmt.Println("Moves:")
	fmt.Println(&mlist)

	b.SetFEN("r2qkb1r/pp2p1Pp/1np1bn2/3p4/3P4/8/PPP2PP1/RNBQKBNR w KQkq - 0 9")
	fmt.Println(&b)
	mlist.Clean()
	b.GeneratePawnMoves(&mlist, b.Player)

	fmt.Println("Moves:")
	fmt.Println(&mlist)

	b.SetFEN("r2qkbnr/pp2p1Pp/1np1b3/4P3/1PpP1P2/8/P6p/RNBQKBN1 b Qkq b3 0 12")
	fmt.Println(&b)
	mlist.Clean()
	b.GeneratePawnMoves(&mlist, b.Player)

	fmt.Println("Moves:")
	fmt.Println(&mlist)

}
