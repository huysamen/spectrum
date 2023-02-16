package spectrum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRGBToRGBui8(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want [3]uint8
	}

	for _, c := range AllColours {
		tests = append(tests, struct {
			name string
			args [3]float64
			want [3]uint8
		}{name: c.Name, args: c.RGB, want: c.RGBui8})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := RGBToRGBui8(tt.args[0], tt.args[1], tt.args[2])

			assert.Equal(t, tt.want[0], r, "RGBToRGBui8(): R = %v, want %v", r, tt.want[0])
			assert.Equal(t, tt.want[1], g, "RGBToRGBui8(): G = %v, want %v", g, tt.want[1])
			assert.Equal(t, tt.want[2], b, "RGBToRGBui8(): B = %v, want %v", b, tt.want[2])
		})
	}
}

func TestRGBui8ToRGB(t *testing.T) {
	var tests []struct {
		name string
		args [3]uint8
		want [3]float64
	}

	for _, c := range AllColours {
		tests = append(tests, struct {
			name string
			args [3]uint8
			want [3]float64
		}{name: c.Name, args: c.RGBui8, want: c.RGB})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := RGBui8ToRGB(tt.args[0], tt.args[1], tt.args[2])

			assert.InDeltaf(t, tt.want[0], r, 0.00000001, "RGBui8ToRGB(): R = %v, want %v", r, tt.want[0])
			assert.InDeltaf(t, tt.want[1], g, 0.00000001, "RGBui8ToRGB(): G = %v, want %v", g, tt.want[1])
			assert.InDeltaf(t, tt.want[2], b, 0.00000001, "RGBui8ToRGB(): B = %v, want %v", b, tt.want[2])
		})
	}
}

func TestRGBToRGBui32(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want uint32
	}

	for _, c := range AllColours {
		tests = append(tests, struct {
			name string
			args [3]float64
			want uint32
		}{name: c.Name, args: c.RGB, want: c.RGBui32})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rgb := RGBToRGBui32(tt.args[0], tt.args[1], tt.args[2])

			assert.Equal(t, tt.want, rgb, "RGBToRGBui32(): R = %v, want %v", rgb, tt.want)
		})
	}
}

func TestRGBui32ToRGB(t *testing.T) {
	var tests []struct {
		name string
		args uint32
		want [3]float64
	}

	for _, c := range AllColours {
		tests = append(tests, struct {
			name string
			args uint32
			want [3]float64
		}{name: c.Name, args: c.RGBui32, want: c.RGB})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := RGBui32ToRGB(tt.args)

			assert.InDeltaf(t, tt.want[0], r, 0.00000001, "RGBui8ToRGB(): R = %v, want %v", r, tt.want[0])
			assert.InDeltaf(t, tt.want[1], g, 0.00000001, "RGBui8ToRGB(): G = %v, want %v", g, tt.want[1])
			assert.InDeltaf(t, tt.want[2], b, 0.00000001, "RGBui8ToRGB(): B = %v, want %v", b, tt.want[2])
		})
	}
}
