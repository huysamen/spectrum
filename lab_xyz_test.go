package spectrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabToXYZ(t *testing.T) {
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
		}{name: c.Name, args: c.Lab, want: c.XYZ})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y, z := LabToXYZ(tt.args[0], tt.args[1], tt.args[2])

			assert.InDeltaf(t, tt.want[0], x, 0.00000001, "LabToXYZ(): X = %v, want %v", x, tt.want[0])
			assert.InDeltaf(t, tt.want[1], y, 0.00000001, "LabToXYZ(): Y = %v, want %v", y, tt.want[1])
			assert.InDeltaf(t, tt.want[2], z, 0.00000001, "LabToXYZ(): Z = %v, want %v", z, tt.want[2])
		})
	}
}
