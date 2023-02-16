package spectrum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCMYKToCMY(t *testing.T) {
	var tests []struct {
		name string
		args [4]float64
		want [3]float64
	}

	for _, c := range AllColours {
		tests = append(tests, struct {
			name string
			args [4]float64
			want [3]float64
		}{name: c.Name, args: c.CMYK, want: c.CMY})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, m, y := CMYKToCMY(tt.args[0], tt.args[1], tt.args[2], tt.args[3])

			assert.InDeltaf(t, tt.want[0], c, 0.000001, "CMYKToCMY(): C = %v, want %v", c, tt.want[0])
			assert.InDeltaf(t, tt.want[1], m, 0.000001, "CMYKToCMY(): M = %v, want %v", m, tt.want[1])
			assert.InDeltaf(t, tt.want[2], y, 0.000001, "CMYKToCMY(): Y = %v, want %v", y, tt.want[2])
		})
	}
}
