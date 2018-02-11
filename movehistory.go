package chesskimo

// TODO make this variable in size

const (
	max_history_moves = 1024
)

type MoveHistory struct {
	size  uint16
	moves [max_history_moves]Move
}

func (mh *MoveHistory) Reset() {
	mh.size = 0
}

func (mh *MoveHistory) Put(m Move) {
	// TODO -- should we handle excess of 512 moves?
	mh.moves[mh.size] = m
	mh.size++
}

func (mh *MoveHistory) Get(n uint16) Move {
	return mh.moves[n]
}
