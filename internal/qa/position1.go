package qa

type Position1 int

const (
	InsideBoth Position1 = iota
	InsideRectangle
	OutsideBoth
)

func (p Position1) String() string {
	switch p {
	case InsideBoth:
		return "InsideBoth"
	case InsideRectangle:
		return "InsideRectangle"
	case OutsideBoth:
		return "OutsideBoth"
	default:
		return "Unknown"
	}
}
