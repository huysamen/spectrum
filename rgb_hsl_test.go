package spectrum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRGBToHSL(t *testing.T) {
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
		}{name: c.Name, args: c.RGB, want: c.HSL})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, s, l := RGBToHSL(tt.args[0], tt.args[1], tt.args[2])

			assert.InDeltaf(t, tt.want[0], h, 0.00000001, "RGBToHSL(): H = %v, want %v", h, tt.want[0])
			assert.InDeltaf(t, tt.want[1], s, 0.00000001, "RGBToHSL(): S = %v, want %v", s, tt.want[1])
			assert.InDeltaf(t, tt.want[2], l, 0.00000001, "RGBToHSL(): L = %v, want %v", l, tt.want[2])
		})
	}
}
