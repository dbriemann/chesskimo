package main

import (
	"fmt"

	"github.com/dbriemann/chesskimo/base"
	"github.com/dbriemann/chesskimo/engine"
)

func main() {
	fens := []string{
		//		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		//		"rnbqkb1r/pp2pppp/5n2/2ppP3/3P4/8/PPP2PPP/RNBQKBNR w KQkq d6 0 4",
		//		"r2qkbnr/pp2p1Pp/1np1b3/3pPp2/3P4/8/PPP2PP1/RNBQKBNR w KQkq f6 0 8",
		//		"r2qkb1r/pp2p1Pp/1np1bn2/3p4/3P4/8/PPP2PP1/RNBQKBNR w KQkq - 0 9",
		//		"r2qkbnr/pp2p1Pp/1np1b3/4P3/1PpP1P2/8/P6p/RNBQKBN1 b Qkq b3 0 12",
		//		"r1bqk2r/ppp2ppp/2n2n2/1B1pp3/1b1PP1P1/2N2N2/PPP2P1P/R1BQK2R b KQkq g3 0 6",
		//		"r3k2r/pppq1ppp/2npbn2/2b1p3/2B1P3/2NPBN2/PPPQ1PPP/R3K2R w KQkq - 4 8",
		//		"8/5k2/2p5/2R5/2p1b3/2K5/1r2R3/1R3n2 w - - 0 1",
		//		"7k/1QQr4/1nrp4/2Qp4/Qn2pp2/3Q1p2/1K4p1/3qQ3 w - - 0 1",
		//		"r1bqk2r/pppp1pp1/2n2n1p/1B2p3/1b2P3/2NP1N2/PPP2PPP/R1BQK2R w KQkq - 2 6", // PIN
		//		"rnb1kb1r/ppppn1pp/4p3/2K1Ppq1/8/8/PPPP1PPP/RNBQ1BNR w kq f6 0 7", // EP PIN - CRAZY STUFF
		//		"r1b1kb1r/ppppn1pp/2n1p3/2K1Ppq1/8/8/PPPPNPPP/RNBQ1B1R w kq - 2 8", // same but no EP
		//		"1k2r3/1pp5/7b/2q5/3PNB2/qr1PKB1r/8/6q1 w - - 0 1", // lots of pins + check
		//		"1k2r3/1pp5/7b/2q5/3PNBN1/qr1PKB1r/4N2B/5Rq1 w - - 0 1", // same with check capture pieces and obstruction pieces
		//		"1k2r3/1pp5/7b/2q5/3PNBn1/qr1PKB1r/8/6q1 w - - 0 1", // lots of pins + double check

		//		"rnb1kbnr/pppp1ppp/8/2q1p3/3PP3/4K3/PPP2PPP/RNBQ1BNR w kq - 5 5",
		"rnb1k2r/pppp1pbp/5np1/2qPp3/1P2P3/4K3/P1P2PPP/RNBQ1BNR w kq - 1 7",
	}

	b := engine.NewBoard()
	mlist := base.MoveList{}

	for _, fen := range fens {
		b.SetFEN(fen)
		fmt.Println(&b)
		mlist.Clear()
		//		b.GenerateAllLegalMoves(&mlist, b.Player)

		//		b.GenerateKingMoves(&mlist, b.Player)
		//		b.GenerateQueenMoves(&mlist, b.Player)
		//		b.GenerateRookMoves(&mlist, b.Player)
		//		b.GenerateBishopMoves(&mlist, b.Player)
		//		b.GenerateKnightMoves(&mlist, b.Player)
		b.GeneratePawnMoves(&mlist, b.Player)

		fmt.Println("Moves:")
		fmt.Println(&mlist)
	}

}
