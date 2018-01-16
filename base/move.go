package base

type Move uint32

const (
	FROM_SHIFT      = 0
	TO_SHIFT        = 8
	PIECE_SHIFT     = 16
	PROMOTION_SHIFT = 20
	EP_SHIFT        = 24
	CAPTURE_SHIFT   = 26
	CASTLE_SHIFT    = 27
	VALUE_SHIFT     = 28
)

func NewMove(from, to, ptype, capflag, promtype, eptype, castleflag, value Move) Move {
	m := from
	to <<= TO_SHIFT
	ptype <<= PIECE_SHIFT
	promtype <<= PROMOTION_SHIFT
	eptype <<= EP_SHIFT
	capflag <<= CAPTURE_SHIFT
	castleflag <<= CASTLE_SHIFT
	value <<= VALUE_SHIFT
	m |= (to | ptype | promtype | eptype | capflag | castleflag | value)

	return m
}

func (m *Move) SetFeature(mask, shift, bits Move) {
	*m &= ^mask
	bits <<= shift
	*m |= bits
}

func (m *Move) GetFeature(mask, shift Move) Move {
	return (*m & mask) >> shift
}

func (m Move) String() string {
	str := ""
	// TODO
	return str
}
