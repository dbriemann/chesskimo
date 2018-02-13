package chesskimo

// PieceList holds squares for a specific set of pieces. This could be all white bishops
// or all black sliders. The maximum size of 16 allows to even hold all pieces of one color.
// This may be a 'waste' of memory in most cases but should be more efficient than allocating
// memory dynamically.
type PieceList struct {
	Size   uint8
	Pieces [16]Square
}

func NewPieceList() PieceList {
	p := PieceList{
		Size:   0,
		Pieces: [16]Square{OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB},
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
	p.Size--
	// If the deleted element is not the last, swap it with the last one.
	if idx < p.Size {
		p.Pieces[idx], p.Pieces[p.Size] = p.Pieces[p.Size], p.Pieces[idx]
	}
}

func (p *PieceList) Clear() {
	p.Size = 0
}
