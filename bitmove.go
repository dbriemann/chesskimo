package chesskimo

// BitMove is structured as follows:
//
// MSB                      LSB
// ... |17 (4 bits) 14||13 (7 bits) 7||6 (7 bits) 0|
// ... |-----prom-----||-----to------||----from----|
// A few bits are 'wasted' here. Squares could be encoded with 6 bits and
// the promotion piece type could be encoded by two bits. Doing it this way
// however allows for easier usage and better performance in this engine.
// The from and to squares are actual 0x88 indexes (0-127) and the promotion
// field contains 1 bit for each possible type: knight, bishop, rook, queen.
//
// TODO this type should be enhanced by an additional value field in the future
// which allows better move ordering in the move search.
type BitMove uint32

const (
	MOVE_FROM_SHIFT      = 0
	MOVE_FROM_MASK       = 0x0000007F
	MOVE_TO_SHIFT        = 7 + 0
	MOVE_TO_MASK         = 0x00003F80
	MOVE_PROMOTION_SHIFT = -2 + 7 + 7 + 0 // We skip 2 bits in the shift because the 2 lowest piece bits are 0 and ignored.
	MOVE_PROMOTION_MASK  = 0x0003C000
)

func NewBitMove(from, to Square, promo Piece) BitMove {
	return (BitMove(promo)<<MOVE_PROMOTION_SHIFT)&MOVE_PROMOTION_MASK | BitMove(to)<<MOVE_TO_SHIFT | BitMove(from)
}

func (m *BitMove) SetFeature(mask, shift BitMove, bits Square) {
	*m &= ^mask                  // Erase bits in the mask's region.
	bits <<= shift               // Move bits to the right place.
	*m |= (BitMove(bits) & mask) // Add them to the move by oring.
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
	p := (m & MOVE_PROMOTION_MASK) >> MOVE_PROMOTION_SHIFT
	return Piece(p)
}

func (m BitMove) All() (Square, Square, Piece) {
	return m.From(), m.To(), m.PromotedPiece()
}

func (m BitMove) MiniNotation() string {
	from := m.From()
	to := m.To()
	promo := m.PromotedPiece()
	str := PrintBoardIndex[from] + PrintBoardIndex[to]
	if promo > 0 {
		str += PrintMap[promo]
	}

	return str
}

//func (m *BitMove) String() string {
//	// TODO - this should be somewhat performant in the end
//	// byte buffers?!
//	from := m.From()
//	to := m.To()
//	prom := m.PromotedPiece()

//	shortCastle := (from == CASTLING_DETECT_SHORT[BLACK][0]) && (to == CASTLE_TYPE_SHORT[BLACK][1]) ||
//		(from == CASTLING_DETECT_SHORT[WHITE][0]) && (to == CASTLE_TYPE_SHORT[WHITE][1])
//	if shortCastle {
//		return "0-0"
//	}

//	longCastle := (from == CASTLING_DETECT_LONG[BLACK][0]) && (to == CASTLING_DETECT_LONG[BLACK][1]) ||
//		(from == CASTLING_DETECT_LONG[WHITE][0]) && (to == CASTLING_DETECT_LONG[WHITE][1])
//	if longCastle {
//		return "0-0-0"
//	}

//	//	str := ""
//	//	//	if !m.PieceType.IsType(PAWN) {
//	//	str += PrintMap[m.Piece]
//	//	//	}
//	//	str += PrintBoardIndex[m.From]
//	//	if m.CapturedPiece != EMPTY {
//	//		str += "x"
//	//		//		if !m.CaptureType.IsType(PAWN) {
//	//		str += PrintMap[m.CapturedPiece]
//	//		//		}
//	//	} else {
//	//		str += "-"
//	//	}
//	//	str += PrintBoardIndex[m.To]
//	//	if m.EpType == EP_TYPE_CAPTURE {
//	//		str += " e.p."
//	//	} else if m.PromotionPiece != EMPTY {
//	//		str += "=" + PrintMap[m.PromotionPiece]
//	//	}
//	//	return str
//}
