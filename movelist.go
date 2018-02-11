package chesskimo

const (
	// If this is not enough at any point the engine will blow :)
	// I could not find any position with that many legal moves
	// so we should be save.
	max_movelist_size = 256
)

type MoveList struct {
	Size  uint16
	Moves [256]Move
}

func (ml *MoveList) Clear() {
	ml.Size = 0
}

func (ml *MoveList) Put(m Move) {
	ml.Moves[ml.Size] = m
	ml.Size++
}

func (ml *MoveList) Get(n uint16) Move {
	return ml.Moves[n]
}

func (ml *MoveList) String() string {
	last := ml.Size - 1
	str := "["
	for i := uint16(0); i < ml.Size; i++ {
		str += ml.Moves[i].String()
		if i < last {
			str += ", "
		}
	}
	str += "]"
	return str
}
