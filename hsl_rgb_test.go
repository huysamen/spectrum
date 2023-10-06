package spectrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHSLToRGB(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want [3]float64
	}

	for _, c := range AllColours {
		tests = append(tests, struct {
			name string
			args [3]float64
			want [3]float64
		}{name: c.Name, args: c.HSL, want: c.RGB})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := HSLToRGB(tt.args[0], tt.args[1], tt.args[2])

			assert.InDeltaf(t, tt.want[0], r, 0.00000001, "HSLToRGB(): R = %v, want %v", r, tt.want[0])
			assert.InDeltaf(t, tt.want[1], g, 0.00000001, "HSLToRGB(): G = %v, want %v", g, tt.want[1])
			assert.InDeltaf(t, tt.want[2], b, 0.00000001, "HSLToRGB(): B = %v, want %v", b, tt.want[2])
		})
	}
}
