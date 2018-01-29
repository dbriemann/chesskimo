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

	b.GeneratePawnMoves(&mlist, base.WHITE)

	fmt.Println("Moves:")
	fmt.Println(&mlist)

	b.SetFEN("rnbqkb1r/pp2pppp/5n2/2ppP3/3P4/8/PPP2PPP/RNBQKBNR w KQkq d6 0 4")
	fmt.Println(&b)
	mlist.Reset()
	b.GeneratePawnMoves(&mlist, base.WHITE)

	fmt.Println("Moves:")
	fmt.Println(&mlist)

	//	b.SetFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	//	fmt.Println()
	//	fmt.Println(&b)

	//	b.SetFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	//	fmt.Println()
	//	fmt.Println(&b)

	//	b.SetFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	//	fmt.Println()
	//	fmt.Println(&b)
}
