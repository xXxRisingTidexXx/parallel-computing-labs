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
