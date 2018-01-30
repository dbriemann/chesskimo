package base

const (
	EP_TYPE_NONE = iota
	EP_TYPE_CAPTURE
	EP_TYPE_CREATE
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
	// TODO - this should be somewhat performant in the end
	// byte buffers?!
	str := ""
	if !m.PieceType.IsType(PAWN) {
		str += PrintMap[m.PieceType]
	}
	str += PrintBoardIndex[m.From]
	if m.CaptureType != NO_PIECE {
		str += "x"
		if !m.CaptureType.IsType(PAWN) {
			str += PrintMap[m.CaptureType]
		}
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
