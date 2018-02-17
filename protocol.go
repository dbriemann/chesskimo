package chesskimo

// Protocol defines how the chess engine talks to the GUI (or other frontends).
type Protocol interface {
	RunInputOutputLoop(engine *Engine)
}
