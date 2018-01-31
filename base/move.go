package base

const (
	EP_TYPE_NONE = iota
	EP_TYPE_CAPTURE
	EP_TYPE_CREATE

	CASTLE_TYPE_NONE = iota
	CASTLE_TYPE_SHORT
	CASTLE_TYPE_LONG
)

type Move struct {
	From          Square
	To            Square
	PieceType     Piece
	CaptureType   Piece
	PromotionType Piece
	EpType        uint8
	CastleType    uint8
}

func NewMove(from, to Square, ptype, captype, promtype Piece, eptype, castletype uint8) Move {
	return Move{
		From:          from,
		To:            to,
		PieceType:     ptype,
		CaptureType:   captype,
		PromotionType: promtype,
		EpType:        eptype,
		CastleType:    castletype,
	}
}

func (m *Move) String() string {
	if m.CastleType == CASTLE_TYPE_SHORT {
		return "0-0"
	} else if m.CastleType == CASTLE_TYPE_LONG {
		return "0-0-0"
	}
	// TODO - this should be somewhat performant in the end
	// byte buffers?!
	str := ""
	//	if !m.PieceType.IsType(PAWN) {
	str += PrintMap[m.PieceType]
	//	}
	str += PrintBoardIndex[m.From]
	if m.CaptureType != NO_PIECE {
		str += "x"
		//		if !m.CaptureType.IsType(PAWN) {
		str += PrintMap[m.CaptureType]
		//		}
	} else {
		str += "-"
	}
	str += PrintBoardIndex[m.To]
	if m.EpType == EP_TYPE_CAPTURE {
		str += " e.p."
	} else if m.PromotionType != NO_PIECE {
		str += "=" + PrintMap[m.PromotionType]
	}
	return str
}
