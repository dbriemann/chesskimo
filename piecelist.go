package chesskimo

// PieceList holds squares for all equal piece types of a single color.
// The size is the maximum of occurences. Some pieces can only occur
// less often but we use 10 as maximum for all types (and waste a few
// bytes), so we have a consistent memory usage.
type PieceList struct {
	Size   uint8
	Pieces [10]Square
}

func NewPieceList() PieceList {
	p := PieceList{
		Size:   0,
		Pieces: [10]Square{OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB},
	}
	return p
}

func (p *PieceList) Add(sq Square) {
	p.Pieces[p.Size] = sq
	p.Size++
}

func (p *PieceList) Move(from, to Square) {
	for i := uint8(0); i < p.Size; i++ {
		if p.Pieces[i] == from {
			p.Pieces[i] = to
		}
	}
}

func (p *PieceList) Remove(sq Square) {
	for i := uint8(0); i < p.Size; i++ {
		if p.Pieces[i] == sq {
			p.RemoveIdx(i)
		}
	}
}

func (p *PieceList) RemoveIdx(idx uint8) {
	p.Pieces[idx] = OTB // TODO remove, currently kept for testing
	p.Size--
	// If the deleted element is not the last, swap it with the last one.
	if idx < p.Size {
		p.Pieces[idx], p.Pieces[p.Size] = p.Pieces[p.Size], p.Pieces[idx]
	}
}

func (p *PieceList) Clear() {
	// TODO remove OTB stuff.
	for i := 0; i < 10; i++ {
		p.Pieces[i] = OTB
	}

	p.Size = 0
}
