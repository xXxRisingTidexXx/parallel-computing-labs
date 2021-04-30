package qa_test

import (
	"github.com/xXxRisingTidexXx/parallel-computing-labs/internal/qa"
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
			t.Errorf("qa_test: f(%f, %f) caused %v", s.x, s.y, p)
		}
	}()
	if a := qa.Compute(s.x, s.y); a != s.a {
		t.Errorf("qa_test: f(%f, %f), %.12f != %.12f", s.x, s.y, a, s.a)
	}
}

func testNegativeCompute(t *testing.T, s computeSpec) {
	defer func() {
		if recover() == nil {
			t.Errorf("qa_test: f(%f, %f) caused no panic", s.x, s.y)
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
	for _, spec := range specs {
		if count := qa.DescentCount(spec.a); count != spec.count {
			t.Errorf("qa_test: f(%v), %d != %d", spec.a, count, spec.count)
		}
	}
}

func TestRecognizeShape(t *testing.T) {
	specs := []recognizeShapeSpec{
		{f: 18, shape: qa.Inconsistency},
		{a: -4.3, b: 4.4, c: -2.7, shape: qa.IntersectingLines},
		{3.6, 0.4, -2.7, 6.8, -7, 3.3, qa.Hyperbola},
		{shape: qa.ParallelLines},
		{a: 2.8, f: -3, shape: qa.ParallelLines},
		{a: 2.9, d: 3.6, e: -8.4, f: 0.9, shape: qa.Parabola},
		{2.9, 0.7, 1.4, -5.4, -3.6, -6.8, qa.Ellipse},
		{a: 2.9, b: 0.7, c: 1.4, shape: qa.ImaginaryLines},
	}
	for _, spec := range specs {
		shape := qa.RecognizeShape(spec.a, spec.b, spec.c, spec.d, spec.e, spec.f)
		if shape != spec.shape {
			t.Errorf(
				"qa_test: f(%.1f, %.1f, %.1f, %.1f, %.1f, %.1f), %s != %s",
				spec.a,
				spec.b,
				spec.c,
				spec.d,
				spec.e,
				spec.f,
				shape,
				spec.shape,
			)
		}
	}
}
