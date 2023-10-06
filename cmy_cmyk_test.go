package spectrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCMYToCMYK(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want [4]float64
	}

	for _, c := range AllColours {
		tests = append(tests, struct {
			name string
			args [3]float64
			want [4]float64
		}{name: c.Name, args: c.CMY, want: c.CMYK})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, m, y, k := CMYToCMYK(tt.args[0], tt.args[1], tt.args[2])

			assert.InDeltaf(t, tt.want[0], c, 0.00000001, "CMYToCMYK(): C = %v, want %v", c, tt.want[0])
			assert.InDeltaf(t, tt.want[1], m, 0.00000001, "CMYToCMYK(): M = %v, want %v", m, tt.want[1])
			assert.InDeltaf(t, tt.want[2], y, 0.00000001, "CMYToCMYK(): Y = %v, want %v", y, tt.want[2])
			assert.InDeltaf(t, tt.want[3], k, 0.00000001, "CMYToCMYK(): K = %v, want %v", k, tt.want[3])
		})
	}
}
