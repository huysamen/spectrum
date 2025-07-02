package spectrum

import (
	"math"
	"testing"
)

func almostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestRGBToLab(t *testing.T) {
	tests := []struct {
		r, g, b  float64
		l, a, b2 float64
	}{
		// White
		{1, 1, 1, 100, 0, 0},
		// Black
		{0, 0, 0, 0, 0, 0},
		// Red
		{1, 0, 0, 53.24, 80.09, 67.20},
		// Green
		{0, 1, 0, 87.74, -86.18, 83.18},
		// Blue
		{0, 0, 1, 32.30, 79.19, -107.86},
	}
	tol := 0.2 // Allowable tolerance due to floating point math
	for _, tt := range tests {
		l, a, b2 := RGBToLab(tt.r, tt.g, tt.b)

		if !almostEqual(l, tt.l, tol) || !almostEqual(a, tt.a, tol) || !almostEqual(b2, tt.b2, tol) {
			t.Errorf("RGBToLab(%v, %v, %v) = (%v, %v, %v), want (%v, %v, %v)", tt.r, tt.g, tt.b, l, a, b2, tt.l, tt.a, tt.b2)
		}
	}
}
