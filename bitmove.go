package chesskimo

// BitMove is structured as follows.
//
// MSB                      LSB
// ... |17  14||13   7||6    0|
// ... |-prom-||--to--||-from-|
type BitMove uint32

const (
	MOVE_FROM_SHIFT          = 0
	MOVE_FROM_MASK           = 0x0000007F // 0-127 squares because of 0x88 board structure.
	MOVE_TO_SHIFT            = 7 + 0
	MOVE_TO_MASK             = 0x00003F80 // 0-127
	MOVE_PROMOTION_SHIFT     = 4 + 7 + 0
	MOVE_PROMOTION_SHIFT_OUT = 4 - 2 + 7 + 0 // To adhere to pieces encoding.
	MOVE_PROMOTION_MASK      = 0x0003C000

	// TODO -- value for move ordering
)

func NewBitMove(from, to, promo BitMove) BitMove {
	return promo<<MOVE_PROMOTION_SHIFT | to<<MOVE_TO_SHIFT | from
}

func (m *BitMove) SetFeature(mask, shift, bits BitMove) {
	*m &= ^mask    // Erase bits in the mask's region.
	bits <<= shift // Move bits to the right place.
	*m |= bits     // Add them to the move by oring.
}

func (m BitMove) From() Square {
	sq := m & MOVE_FROM_MASK // Needs no shift.
	return Square(sq)
}

func (m BitMove) To() Square {
	sq := (m & MOVE_TO_MASK) >> MOVE_TO_SHIFT
	return Square(sq)
}

func (m BitMove) PromotedPiece() Piece {
	p := (m & MOVE_PROMOTION_MASK) >> MOVE_PROMOTION_SHIFT_OUT
	return Piece(p)
}

func (m BitMove) MiniNotation() string {
	from := m.From()
	to := m.To()
	promo := m.PromotedPiece()
	str := PrintBoardIndex[from] + PrintBoardIndex[to]
	if promo > 0 {
		ptype := Piece(promo << 2)
		str += PrintMap[ptype]
	}

	return str
}

func (m BitMove) String() string {
	str := ""
	// TODO
	return str
}
