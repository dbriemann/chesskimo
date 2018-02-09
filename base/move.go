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
	From           Square
	To             Square
	Piece          Piece
	CapturedPiece  Piece
	PromotionPiece Piece
	EpType         uint8
	CastleType     uint8
}

func NewMove(from, to Square, piece, cappiece, prompiece Piece, eptype, castletype uint8) Move {
	return Move{
		From:           from,
		To:             to,
		Piece:          piece,
		CapturedPiece:  cappiece,
		PromotionPiece: prompiece,
		EpType:         eptype,
		CastleType:     castletype,
	}
}

func (m *Move) String() string {
	// TODO - this should be somewhat performant in the end
	// byte buffers?!

	if m.CastleType == CASTLE_TYPE_SHORT {
		return "0-0"
	} else if m.CastleType == CASTLE_TYPE_LONG {
		return "0-0-0"
	}
	str := ""
	//	if !m.PieceType.IsType(PAWN) {
	str += PrintMap[m.Piece]
	//	}
	str += PrintBoardIndex[m.From]
	if m.CapturedPiece != EMPTY {
		str += "x"
		//		if !m.CaptureType.IsType(PAWN) {
		str += PrintMap[m.CapturedPiece]
		//		}
	} else {
		str += "-"
	}
	str += PrintBoardIndex[m.To]
	if m.EpType == EP_TYPE_CAPTURE {
		str += " e.p."
	} else if m.PromotionPiece != EMPTY {
		str += "=" + PrintMap[m.PromotionPiece]
	}
	return str
}
