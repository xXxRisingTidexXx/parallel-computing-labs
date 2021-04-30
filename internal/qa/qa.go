package qa

func Compute(x, y float64) float64 {
	d := 2*x - x*x*y
	if d == 0 {
		panic("qa: got 0 in denominator")
	}
	return (x*x + 2*x*y - y*y) / d
}

func DescentCount(a []float64) int { // 1
	var count int                 // 2
	for i := 0; i < len(a); i++ { // 3
		if i == len(a)-1 || a[i] <= a[i+1] { // 4, 5
			count++ // 6
		}
	}
	return count // 7
}

func RecognizeShape(a, b, c, d, e, f float64) Shape {
	small := a*c - b*b
	isComposite := a*c*f + 2*b*e*d - d*c*d - b*b*f - a*e*e == 0
	if small < 0 {
		if isComposite {
			return IntersectingLines
		}
		return Hyperbola
	}
	if small == 0 {
		if isComposite {
			return ParallelLines
		}
		return Parabola
	}
	if isComposite {
		return ImaginaryLines
	}
	return Ellipse
}

func RecognizePosition1() Position1 {
	return OutsideBoth
}

func RecognizePosition2() Position2 {
	return OutsideTriangle
}
