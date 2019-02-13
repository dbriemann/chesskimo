package chesskimo

// SearchResult contains all relevant info that should
// be returned from a best move search.
type SearchResult struct {
	Move  BitMove
	Score int
	Depth int
}

// SearchSettings defines constraints that may exist for
// the search.
type SearchSettings struct {
	MaxDepth int
}

// SearchFun function type defines how a search function
// must be defined.
type SearchFun func(*Engine, *SearchSettings, *uint32) SearchResult

// Communicator defines how the chess engine talks to the GUI (or other frontends).
type Communicator interface {
	RunInputOutputLoop(engine *Engine)
}
