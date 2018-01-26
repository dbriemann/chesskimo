package base

// PieceList holds squares for all equal piece types of a single color.
type PieceList []Square

func NewPieceList(size int) PieceList {
	return make([]Square, size)
}

func (p PieceList) Add(sq Square) {
	for i := 0; i < len(p); i++ {
		if p[i] == OTB {
			p[i] = sq
		}
	}
}

func (p PieceList) Remove(sq Square) {
	for i := 0; i < len(p); i++ {
		if p[i] == sq {
			p[i] = OTB
		}
	}
}

func (p PieceList) RemoveIdx(idx int) {
	p[idx] = OTB
}
