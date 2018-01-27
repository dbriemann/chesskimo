package base

const (
	max_moves = 512
)

type MoveHistory struct {
	size  uint16
	moves [max_moves]Move
}

func (ml *MoveHistory) Reset() {
	ml.size = 0
}

func (ml *MoveHistory) Put(m Move) {
	// TODO -- should we handle excess of 512 moves?
	ml.moves[ml.size] = m
	ml.size++
}

func (ml *MoveHistory) Get(n uint16) Move {
	return ml.moves[n]
}
