package qa

type Position2 int

const (
	InsideTriangle Position2 = iota
	OutsideTriangle
)

func (p Position2) String() string {
	switch p {
	case InsideTriangle:
		return "InsideTriangle"
	case OutsideTriangle:
		return "OutsideTriangle"
	default:
		return "Unknown"
	}
}
