package chesskimo

import "time"

// AlphaBeta runs a minmaxing search with alpha-beta pruning for the given parameters.
// It returns the best move for a full search of a given max depth. If a depth search has to
// be cancelled, the result of the next best depth is returned.
func AlphaBeta(engine *Engine, ss *SearchSettings, dostop *uint32) SearchResult {
	startTime := time.Now()
	sr := SearchResult{Move: BitMove(0)}

	sr.Depth = ss.MaxDepth
	valid := alphaBeta(engine, ss.MaxDepth, -INFINITY, INFINITY, dostop, &startTime, &sr)
	if valid {

	}

	return sr
}

func alphaBeta(engine *Engine, depth, alpha, beta int, dostop *uint32, startTime *time.Time, result *SearchResult) bool {
	mlist := MoveList{}
	board := engine.board

	return true
}
