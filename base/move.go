package base

type Move struct {
	From Square
	To   Square
}

func (m *Move) String() string {
	// TODO - this should be performant in the end
	str := ""
	return str
}
