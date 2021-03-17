package qa

func Compute(x, y float64) float64 {
	d := 2 * x - x * x * y
	if d == 0 {
		panic("qa: got 0 in denominator")
	}
	return (x * x + 2 * x * y - y * y) / d
}
