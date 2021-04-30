package qa

type Shape int

const (
	Inconsistency Shape = iota
	ImaginaryLines
	ParallelLines
	IntersectingLines
	Hyperbola
	Parabola
	Ellipse
)

func (s Shape) String() string {
	switch s {
	case Inconsistency:
		return "Inconsistency"
	case ImaginaryLines:
		return "ImaginaryLines"
	case ParallelLines:
		return "ParallelLines"
	case IntersectingLines:
		return "IntersectingLines"
	case Hyperbola:
		return "Hyperbola"
	case Parabola:
		return "Parabola"
	case Ellipse:
		return "Ellipse"
	default:
		return "Unknown"
	}
}
