package chesskimo

import (
	"math/rand"
	"time"
)

// SimpleMCSearch runs a simple random simulation (monte carlo) for a
// given time and returns the best move.
func SimpleMCSearch(engine *Engine, ts *TimeSettings, dostop *uint32) SearchResult {
	startTime := time.Now()
	board := &engine.board
	sr := SearchResult{BestMove: BitMove(0)}
	player := board.Player
	mlist := MoveList{}
	workMlist := MoveList{}
	workBoard := *board
	simcount := uint64(0)

	// Find all possible first moves.
	board.GenerateAllLegalMoves(&mlist)
	scores := make([]int64, mlist.Size)

	maxtime := 10.0

	for /*atomic.LoadUint32(dostop) == 0 &&*/ time.Since(startTime).Seconds() < maxtime {
		for i := uint32(0); i < mlist.Size; i++ {
			move := mlist.Moves[i]
			//			engine.logger.Println("Simulation for move ", move.MiniNotation())
			workBoard = *board
			workBoard.MakeLegalMove(move)
			for { // run sim until game ends
				//				engine.logger.Println(workBoard.String())
				// 1. Generate moves.
				//				engine.logger.Println("genall start, move", workBoard.MoveNumber)
				workMlist.Clear()
				workBoard.GenerateAllLegalMoves(&workMlist)
				//				engine.logger.Println("genall: ", workMlist.String())
				// e2e4 h7h5 d2d4
				if workBoard.MoveNumber > 150 {
					// artificial limit -> draw
					break
				} else if workMlist.Size == 0 {
					// No moves for current player
					if workBoard.CheckInfo == CHECK_NONE {
						// No check -> stalemate
						break
					} else {
						// Check -> checkmate
						if workBoard.Player != player {
							scores[i] += 1
						} else {
							scores[i] += -1
						}
						break
					}
				}
				// 2. Make random move
				r := rand.Intn(int(workMlist.Size))
				//				engine.logger.Printf("MAKE MOVE %s", workMlist.Moves[r].MiniNotation())
				//				engine.logger.Print(workBoard.InfoBoardString())
				workBoard.MakeLegalMove(workMlist.Moves[r])
			}
			simcount++

			if time.Since(startTime).Seconds() >= maxtime {
				break
			}
		}
	}

	engine.logger.Printf("Time used: %f sec. Simulations run %d.", time.Since(startTime).Seconds(), simcount)

	bestscore := -int64(simcount)
	// Find best move and log data.
	for i := 0; i < len(scores); i++ {
		score := scores[i]
		if score > bestscore {
			sr.BestMove = mlist.Moves[i]
			bestscore = score
		}
		engine.logger.Printf("Move %s has score %d", mlist.Moves[i].MiniNotation(), score)
	}

	return sr
}
