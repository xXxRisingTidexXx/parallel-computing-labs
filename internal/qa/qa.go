package qa

import (
	"encoding/json"
	"math"
	"os"
	"strconv"
	"strings"
)

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
	isComposite := a*c*f+2*b*e*d-d*c*d-b*b*f-a*e*e == 0
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

func RecognizePosition1(x, y float64) Position1 {
	if math.Pow(x-6, 2)+math.Pow(y-3, 2) <= 4 && x <= 6 {
		return InsideBoth
	}
	if 2 <= x && x <= 10 && 0 <= y && y <= 6 {
		return InsideRectangle
	}
	return OutsideBoth
}

func RecognizePosition2(x, y float64) Position2 {
	if 0 <= y && (-2 <= x && x <= 0 && y <= 1.25*x+2.5 || 0 <= x && x <= 2 && y <= -1.25*x+2.5) {
		return InsideTriangle
	}
	return OutsideTriangle
}

func CumulateMeans(a []float64) []float64 {
	b := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		if i == 0 {
			b[i] = a[i]
		} else {
			b[i] = (float64(i)*b[i-1] + a[i]) / (float64(i) + 1)
		}
	}
	return b
}

func FilterEvens(a []float64) []float64 {
	b := make([]float64, 0)
	for i := 0; i < len(a); i++ {
		if math.Mod(a[i], 2) == 0 {
			b = append(b, a[i])
		}
	}
	return b
}

func ParseSlice(s string) ([]float64, error) {
	symbols := strings.Split(s, ", ")
	numbers := make([]float64, len(symbols))
	for i := range symbols {
		number, err := strconv.ParseFloat(symbols[i], 64)
		if err != nil {
			return nil, err
		}
		numbers[i] = number
	}
	return numbers, nil
}

func ReadWorkers(path string) ([]Worker, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var workers []Worker
	if err := json.NewDecoder(file).Decode(&workers); err != nil {
		_ = file.Close()
		return nil, err
	}
	if err := file.Close(); err != nil {
		return nil, err
	}
	return workers, nil
}

func ReadWorkersByOccupation(occupation string) ([]Worker, error) {
	workers, err := ReadWorkers("testdata/workers.json")
	if err != nil {
		return nil, err
	}
	newWorkers := make([]Worker, 0)
	for _, worker := range workers {
		if worker.Occupation == occupation {
			newWorkers = append(newWorkers, worker)
		}
	}
	return newWorkers, nil
}

func ReadWorkersGTSalary(salary float64) ([]Worker, error) {
	workers, err := ReadWorkers("testdata/workers.json")
	if err != nil {
		return nil, err
	}
	newWorkers := make([]Worker, 0)
	for _, worker := range workers {
		if worker.Salary > salary {
			newWorkers = append(newWorkers, worker)
		}
	}
	return newWorkers, nil
}
