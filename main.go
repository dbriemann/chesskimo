package main

import (
	"fmt"

	"github.com/zensword/chesskimo/board"
)

func main() {
	bb := board.NewRootBB()
	fmt.Println()

	// TODO: put this all in tests.

	bb.SetFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	fmt.Println()
	bb.SetFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	fmt.Println()
	bb.SetFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	fmt.Println()

	fmt.Println(bb)
}
