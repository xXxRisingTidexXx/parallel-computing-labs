package qa

func Compute(x, y float64) float64 {
	d := 2*x - x*x*y
	if d == 0 {
		panic("qa: got 0 in denominator")
	}
	return (x*x + 2*x*y - y*y) / d
}

func DescentCount(a []float64) int { // 1
	var count int // 2
	for i := 0; i < len(a); i++ { // 3
		if i == len(a)-1 || a[i] <= a[i+1] { // 4, 5
			count++ // 6
		}
	}
	return count // 7
}
