package chesskimo

// SearchResult contains all relevant info that should
// be returned from a best move search.
type SearchResult struct {
	BestMove BitMove
}

// TimeSettings defines constraints that may exist for
// the search.
type TimeSettings struct {
}

// SearchFun function type defines how a search function
// must be defined.
type SearchFun func(*Engine, *TimeSettings, *uint32) SearchResult

// Communicator defines how the chess engine talks to the GUI (or other frontends).
type Communicator interface {
	RunInputOutputLoop(engine *Engine)
}
