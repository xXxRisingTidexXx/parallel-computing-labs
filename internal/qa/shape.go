package qa

type Shape int

const (
	ImaginaryLines Shape = iota
	ParallelLines
	IntersectingLines
	Hyperbola
	Parabola
	Ellipse
)

func (s Shape) String() string {
	switch s {
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
