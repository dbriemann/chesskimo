package chesskimo

const (
	// If this is not enough at any point the engine will blow :)
	// I could not find any position with that many legal moves
	// so we should be save.
	max_movelist_size = 256
)

type MoveList struct {
	Size  uint32
	Moves [256]BitMove
}

func (ml *MoveList) Clear() {
	ml.Size = 0
}

func (ml *MoveList) Put(m BitMove) {
	ml.Moves[ml.Size] = m
	ml.Size++
}

func (ml *MoveList) Get(n uint16) BitMove {
	return ml.Moves[n]
}

func (ml *MoveList) String() string {
	last := ml.Size - 1
	str := "["
	for i := uint32(0); i < ml.Size; i++ {
		str += ml.Moves[i].MiniNotation()
		if i < last {
			str += ", "
		}
	}
	str += "]"
	return str
}
