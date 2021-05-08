package qa_test

import (
	"github.com/xXxRisingTidexXx/parallel-computing-labs/internal/qa"
	"math"
	"testing"
)

func TestCompute(t *testing.T) {
	positiveSpecs := []computeSpec{
		{2, 3, -0.875},
		{5, 2, -1.025},
		{11, 0, 5.5},
		{-28, 0, -14},
		{1056.34, 0, 528.17},
		{2, -3, -1.0625},
	}
	for _, s := range positiveSpecs {
		testPositiveCompute(t, s)
	}
	negativeSpecs := []computeSpec{
		{0, 14, 0},
		{0, 0, 2},
		{0, -234, 0},
		{0, 2378123, 0},
		{0, 0.3, 4},
		{0, -29.0341, -2.1},
		{1, 2, 1},
		{2, 1, 1},
		{0.5, 4, 0},
		{4, 0.5, 0},
		{-0.3, -20.0 / 3, 5},
		{-0.000001, -2000000, 0.25},
	}
	for _, s := range negativeSpecs {
		testNegativeCompute(t, s)
	}
}

func testPositiveCompute(t *testing.T, s computeSpec) {
	defer func() {
		if p := recover(); p != nil {
			t.Errorf("Compute(%f, %f) caused %v", s.x, s.y, p)
		}
	}()
	if a := qa.Compute(s.x, s.y); a != s.a {
		t.Errorf("Compute(%f, %f), %.12f != %.12f", s.x, s.y, a, s.a)
	}
}

func testNegativeCompute(t *testing.T, s computeSpec) {
	defer func() {
		if recover() == nil {
			t.Errorf("Compute(%f, %f) caused no panic", s.x, s.y)
		}
	}()
	_ = qa.Compute(s.x, s.y)
}

func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = qa.Compute(12.29, -32.28)
	}
}

func TestDescentCount(t *testing.T) {
	specs := []descentCountSpec{
		{[]float64{}, 0},
		{[]float64{2}, 1},
		{[]float64{0}, 1},
		{[]float64{-400}, 1},
		{[]float64{-400, -400}, 2},
		{[]float64{2, 2, 2, 2, 2, 2, 2}, 7},
		{[]float64{1, 2, 3, 4, 5, 6, 7}, 7},
		{[]float64{7, 6, 5, 4, 3, 2, 1}, 1},
		{[]float64{-7, -6, -5, -4, -3, -2, -1}, 7},
		{[]float64{7, 6, 5, 4, 4, 2, 1}, 2},
		{[]float64{7, 6, 5, 4, 4, 4, 1}, 3},
		{[]float64{7, 6, 7, 4, 7, 4, 7}, 4},
	}
	for i, spec := range specs {
		if count := qa.DescentCount(spec.a); count != spec.count {
			t.Errorf("DescentCount[%d], %d != %d", i, count, spec.count)
		}
	}
}

func TestRecognizeShape(t *testing.T) {
	specs := []recognizeShapeSpec{
		{a: -4.3, b: 4.4, c: -2.7, shape: qa.IntersectingLines},
		{3.6, 0.4, -2.7, 6.8, -7, 3.3, qa.Hyperbola},
		{shape: qa.ParallelLines},
		{f: 18, shape: qa.ParallelLines},
		{a: 2.8, f: -3, shape: qa.ParallelLines},
		{a: 2.9, d: 3.6, e: -8.4, f: 0.9, shape: qa.Parabola},
		{2.9, 0.7, 1.4, -5.4, -3.6, -6.8, qa.Ellipse},
		{a: 2.9, b: 0.7, c: 1.4, shape: qa.ImaginaryLines},
	}
	for i, spec := range specs {
		shape := qa.RecognizeShape(spec.a, spec.b, spec.c, spec.d, spec.e, spec.f)
		if shape != spec.shape {
			t.Errorf("RecognizeShape[%d], %s != %s", i, shape, spec.shape)
		}
	}
}

func TestRecognizePosition1(t *testing.T) {
	specs := []recognizePosition1Spec{
		{3, 6, qa.InsideRectangle},
		{2, 6, qa.InsideRectangle},
		{3, 1, qa.InsideRectangle},
		{3, 3, qa.InsideRectangle},
		{1, 1, qa.OutsideBoth},
		{2, 3, qa.InsideRectangle},
		{4, 2, qa.InsideRectangle},
		{0, 2, qa.OutsideBoth},
	}
	for i, spec := range specs {
		if position := qa.RecognizePosition1(spec.x, spec.y); position != spec.position {
			t.Errorf("RecognizePosition1[%d], %s != %s", i, position, spec.position)
		}
	}
}

func TestRecognizePosition2(t *testing.T) {
	specs := []recognizePosition2Spec{
		{2, 0, qa.InsideTriangle},
		{2, 2, qa.OutsideTriangle},
		{4, 3, qa.OutsideTriangle},
		{4, 2, qa.OutsideTriangle},
		{0, 2, qa.InsideTriangle},
		{2, 1, qa.OutsideTriangle},
		{3, 3, qa.OutsideTriangle},
		{2, 4, qa.OutsideTriangle},
	}
	for i, spec := range specs {
		if position := qa.RecognizePosition2(spec.x, spec.y); position != spec.position {
			t.Errorf("RecognizePosition2[%d], %s != %s", i, position, spec.position)
		}
	}
}

func TestCumulateMeans(t *testing.T) {
	specs := []cumulateMeansSpec{
		{},
		{[]float64{2}, []float64{2}},
		{[]float64{2, 54}, []float64{2, 28}},
		{[]float64{-23, -1, 0, 26}, []float64{-23, -12, -8, 0.5}},
		{[]float64{2, 2, 2, 2, 2, 2, 2}, []float64{2, 2, 2, 2, 2, 2, 2}},
		{
			[]float64{2, -2, 2, -2, 2, -2, 2},
			[]float64{2, 0, 0.666666667, 0, 0.4, 0, 0.28571428},
		},
		{
			[]float64{23, 28, 1, 8, 9, 0, 0.2},
			[]float64{23, 25.5, 17.33333333, 15, 13.8, 11.5, 9.885714285},
		},
	}
	for _, spec := range specs {
		testSlices(t, spec.b, qa.CumulateMeans(spec.a))
	}
}

func testSlices(t *testing.T, actual, predicted []float64) {
	if len(actual) != len(predicted) {
		t.Fatalf("Slices lengths, %d != %d", len(actual), len(predicted))
	}
	for i := range predicted {
		if math.Abs(actual[i]-predicted[i]) >= 0.000001 {
			t.Errorf("Slices[%d], %.12f != %.12f", i, actual[i], predicted[i])
		}
	}
}

func TestFilterEvens(t *testing.T) {
	specs := []filterEvensSpec{
		{},
		{[]float64{2}, []float64{2}},
		{[]float64{2, 28}, []float64{2, 28}},
		{[]float64{-23, -12, -8, 0.5}, []float64{-12, -8}},
		{[]float64{2, 2, 2, 2, 2, 2, 2}, []float64{2, 2, 2, 2, 2, 2, 2}},
		{[]float64{2, 0, 0.666666667, 0, 0.4, 0, 0.28571428}, []float64{2, 0, 0, 0}},
		{[]float64{23, 25.5, 17.33333333, 15, 13.8, 11.5, 9.885714285}, []float64{}},
	}
	for _, spec := range specs {
		testSlices(t, spec.b, qa.FilterEvens(spec.a))
	}
}

func TestParseSlice(t *testing.T) {
	specs := []parseSliceSpec{
		{"", true, []float64{}},
		{"1", true, []float64{1}},
		{"-5, 34.5", true, []float64{-5, 34.5}},
		{"0, 0, 0, 0, 0, 0", true, []float64{0, 0, 0, 0, 0, 0}},
		{s: "-2,,,,,45"},
		{s: "-2   45"},
		{s: "-2, 28.asd2383d, 45"},
	}
	for i, spec := range specs {
		a, err := qa.ParseSlice(spec.s)
		if spec.isValid {
			testSlices(t, spec.a, a)
		} else if err == nil {
			t.Fatalf("ParseSlice[%d] got no error", i)
		}
	}
}

func TestIntegration(t *testing.T) {
	specs := []integrationSpec{
		{"", true, []float64{}},
		{"2", true, []float64{2}},
		{"-5, 34.5", true, []float64{}},
		{"0, 0, 0, 0, 0, 0", true, []float64{0, 0, 0, 0, 0, 0}},
		{s: "-2,,,,,45"},
		{s: "-2   45"},
		{s: "-2, 28.asd2383d, 45"},
		{"2, 54", true, []float64{2, 28}},
		{"-23, -1, 0, 26", true, []float64{-12, -8}},
		{
			"2, -2, 2, -2, 2, -2, 2",
			true,
			[]float64{2, 0, 0, 0},
		},
	}
	for i, spec := range specs {
		a, err := qa.ParseSlice(spec.s)
		if spec.isValid {
			testSlices(t, spec.b, qa.FilterEvens(qa.CumulateMeans(a)))
		} else if err == nil {
			t.Fatalf("Integration[%d] got no error", i)
		}
	}
}

func TestReadWorkers(t *testing.T) {
	if _, err := qa.ReadWorkers("something.json"); err == nil {
		t.Fatal("ReadWorkers got no error for something.json")
	}
	if _, err := qa.ReadWorkers("testdata/workers.yaml"); err == nil {
		t.Fatal("ReadWorkers got no error for testdata/workers.yaml")
	}
	predicted, err := qa.ReadWorkers("testdata/workers.json")
	if err != nil {
		t.Fatalf("ReadWorkers got an unexpected error, %v", err)
	}
	testWorkers(
		t,
		[]qa.Worker{
			{"Kadenchuk", "Plumber", 2137.9, 14},
			{"Antoshkin", "Technical writer", 0, 0},
			{"Longbottom", "Retailer", 3058, 3.4},
			{"Elizabeth II", "Queen", 304563.32, 68},
			{"Lazo", "Retailer", 428.2, 0.6},
			{"Che Guevara", "Minister of Industries", 12411.56, 7},
			{"Eko", "Retailer", 2109, 6},
			{"Awaji", "Plumber", 1203.5, 3.3},
		},
		predicted,
	)
}

func testWorkers(t *testing.T, actual, predicted []qa.Worker) {
	if len(actual) != len(predicted) {
		t.Fatalf("Workers lengths, %d != %d", len(actual), len(predicted))
	}
	for i := range actual {
		if actual[i] != predicted[i] {
			t.Errorf("Workers[%d], %v != %v", i, actual[i], predicted[i])
		}
	}
}

func TestReadWorkersByOccupation(t *testing.T) {
	specs := []readWorkersByOccupationSpec{
		{},
		{occupation: "plumber"},
		{"Queen", []qa.Worker{{"Elizabeth II", "Queen", 304563.32, 68}}},
		{
			"Retailer",
			[]qa.Worker{
				{"Longbottom", "Retailer", 3058, 3.4},
				{"Lazo", "Retailer", 428.2, 0.6},
				{"Eko", "Retailer", 2109, 6},
			},
		},
	}
	for i, spec := range specs {
		workers, err := qa.ReadWorkersByOccupation(spec.occupation)
		if err != nil {
			t.Fatalf("ReadWorkersByOccupation[%d] got an unexpected error, %v", i, err)
		}
		testWorkers(t, spec.workers, workers)
	}
}

func TestReadWorkersGTSalary(t *testing.T) {
	specs := []readWorkersGTSalarySpec{
		{
			workers: []qa.Worker{
				{"Kadenchuk", "Plumber", 2137.9, 14},
				{"Longbottom", "Retailer", 3058, 3.4},
				{"Elizabeth II", "Queen", 304563.32, 68},
				{"Lazo", "Retailer", 428.2, 0.6},
				{"Che Guevara", "Minister of Industries", 12411.56, 7},
				{"Eko", "Retailer", 2109, 6},
				{"Awaji", "Plumber", 1203.5, 3.3},
			},
		},
		{
			-5000,
			[]qa.Worker{
				{"Kadenchuk", "Plumber", 2137.9, 14},
				{"Antoshkin", "Technical writer", 0, 0},
				{"Longbottom", "Retailer", 3058, 3.4},
				{"Elizabeth II", "Queen", 304563.32, 68},
				{"Lazo", "Retailer", 428.2, 0.6},
				{"Che Guevara", "Minister of Industries", 12411.56, 7},
				{"Eko", "Retailer", 2109, 6},
				{"Awaji", "Plumber", 1203.5, 3.3},
			},
		},
		{salary: 500000},
		{
			2000,
			[]qa.Worker{
				{"Kadenchuk", "Plumber", 2137.9, 14},
				{"Longbottom", "Retailer", 3058, 3.4},
				{"Elizabeth II", "Queen", 304563.32, 68},
				{"Che Guevara", "Minister of Industries", 12411.56, 7},
				{"Eko", "Retailer", 2109, 6},
			},
		},
		{300000, []qa.Worker{{"Elizabeth II", "Queen", 304563.32, 68}}},
	}
	for i, spec := range specs {
		workers, err := qa.ReadWorkersGTSalary(spec.salary)
		if err != nil {
			t.Fatalf("ReadWorkersGTSalary[%d] got an unexpected error, %v", i, err)
		}
		testWorkers(t, spec.workers, workers)
	}
}
