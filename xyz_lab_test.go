package spectrum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXYZToLab(t *testing.T) {
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
		}{name: c.Name, args: c.XYZ, want: c.Lab})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, a, b := XYZToLab(tt.args[0], tt.args[1], tt.args[2])

			assert.InDeltaf(t, tt.want[0], l, 0.00000001, "XYZToLab(): L* = %v, want %v", l, tt.want[0])
			assert.InDeltaf(t, tt.want[1], a, 0.00000001, "XYZToLab(): a* = %v, want %v", a, tt.want[1])
			assert.InDeltaf(t, tt.want[2], b, 0.00000001, "XYZToLab(): b* = %v, want %v", b, tt.want[2])
		})
	}
}
