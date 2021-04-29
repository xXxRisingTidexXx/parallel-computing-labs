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
