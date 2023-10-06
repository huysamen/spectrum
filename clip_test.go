package spectrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClip_RGB(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want [3]float64
	}

	tests = []struct {
		name string
		args [3]float64
		want [3]float64
	}{
		{name: "Under", args: [3]float64{-1.0, -1.0, -1.0}, want: [3]float64{RGBMin, RGBMin, RGBMin}},
		{name: "Normal", args: [3]float64{0.5, 0.5, 0.5}, want: [3]float64{0.5, 0.5, 0.5}},
		{name: "Over", args: [3]float64{2.0, 2.0, 2.0}, want: [3]float64{RGBMax, RGBMax, RGBMax}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := clipRGB(tt.args[0], tt.args[1], tt.args[2])

			assert.Equalf(t, tt.want[0], r, "clipRGB(): R = %v, want %v", r, tt.want[0])
			assert.Equalf(t, tt.want[1], g, "clipRGB(): G = %v, want %v", g, tt.want[1])
			assert.Equalf(t, tt.want[2], b, "clipRGB(): B = %v, want %v", b, tt.want[2])
		})
	}
}

func TestClip_RGBui8(t *testing.T) {
	var tests []struct {
		name string
		args [3]uint8
		want [3]uint8
	}

	tests = []struct {
		name string
		args [3]uint8
		want [3]uint8
	}{
		{name: "Under", args: [3]uint8{0, 0, 0}, want: [3]uint8{RGBui8Min, RGBui8Min, RGBui8Min}},
		{name: "Normal", args: [3]uint8{127, 127, 127}, want: [3]uint8{127, 127, 127}},
		{name: "Over", args: [3]uint8{255, 255, 255}, want: [3]uint8{RGBui8Max, RGBui8Max, RGBui8Max}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := clipRGBui8(tt.args[0], tt.args[1], tt.args[2])

			assert.Equalf(t, tt.want[0], r, "clipRGBui8(): R = %v, want %v", r, tt.want[0])
			assert.Equalf(t, tt.want[1], g, "clipRGBui8(): G = %v, want %v", g, tt.want[1])
			assert.Equalf(t, tt.want[2], b, "clipRGBui8(): B = %v, want %v", b, tt.want[2])
		})
	}
}

func TestClip_XYZ(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want [3]float64
	}

	tests = []struct {
		name string
		args [3]float64
		want [3]float64
	}{
		{name: "Under", args: [3]float64{-1.0, -1.0, -1.0}, want: [3]float64{XYZXMin, XYZYMin, XYZZMin}},
		{name: "Normal", args: [3]float64{0.5, 0.5, 0.5}, want: [3]float64{0.5, 0.5, 0.5}},
		{name: "Over", args: [3]float64{2.0, 2.0, 2.0}, want: [3]float64{XYZXMax, XYZYMax, XYZZMax}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y, z := clipXYZ(tt.args[0], tt.args[1], tt.args[2])

			assert.Equalf(t, tt.want[0], x, "clipXYZ(): X = %v, want %v", x, tt.want[0])
			assert.Equalf(t, tt.want[1], y, "clipXYZ(): Y = %v, want %v", y, tt.want[1])
			assert.Equalf(t, tt.want[2], z, "clipXYZ(): Z = %v, want %v", z, tt.want[2])
		})
	}
}

func TestClip_Lab(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want [3]float64
	}

	tests = []struct {
		name string
		args [3]float64
		want [3]float64
	}{
		{name: "Under", args: [3]float64{-1.0, -129.0, -129.0}, want: [3]float64{LabLMin, LabAMin, LabBMin}},
		{name: "Normal", args: [3]float64{50.0, 0.0, 0.0}, want: [3]float64{50.0, 0.0, 0.0}},
		{name: "Over", args: [3]float64{101.0, 128.0, 128.0}, want: [3]float64{LabLMax, LabAMax, LabBMax}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, a, b := clipLab(tt.args[0], tt.args[1], tt.args[2])

			assert.Equalf(t, tt.want[0], l, "clipLab(): L = %v, want %v", l, tt.want[0])
			assert.Equalf(t, tt.want[1], a, "clipLab(): A = %v, want %v", a, tt.want[1])
			assert.Equalf(t, tt.want[2], b, "clipLab(): B = %v, want %v", b, tt.want[2])
		})
	}
}

func TestClip_HSL(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want [3]float64
	}

	tests = []struct {
		name string
		args [3]float64
		want [3]float64
	}{
		{name: "Under", args: [3]float64{-1.0, -1.0, -1.0}, want: [3]float64{HSLMin, HSLMin, HSLMin}},
		{name: "Normal", args: [3]float64{0.5, 0.5, 0.5}, want: [3]float64{0.5, 0.5, 0.5}},
		{name: "Over", args: [3]float64{2.0, 2.0, 2.0}, want: [3]float64{HSLMax, HSLMax, HSLMax}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, s, l := clipHSL(tt.args[0], tt.args[1], tt.args[2])

			assert.Equalf(t, tt.want[0], h, "clipHSL(): H = %v, want %v", h, tt.want[0])
			assert.Equalf(t, tt.want[1], s, "clipHSL(): S = %v, want %v", s, tt.want[1])
			assert.Equalf(t, tt.want[2], l, "clipHSL(): L = %v, want %v", l, tt.want[2])
		})
	}
}

func TestClip_CMY(t *testing.T) {
	var tests []struct {
		name string
		args [3]float64
		want [3]float64
	}

	tests = []struct {
		name string
		args [3]float64
		want [3]float64
	}{
		{name: "Under", args: [3]float64{-1.0, -1.0, -1.0}, want: [3]float64{CMYMin, CMYMin, CMYMin}},
		{name: "Normal", args: [3]float64{0.5, 0.5, 0.5}, want: [3]float64{0.5, 0.5, 0.5}},
		{name: "Over", args: [3]float64{2.0, 2.0, 2.0}, want: [3]float64{CMYMax, CMYMax, CMYMax}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, m, y := clipCMY(tt.args[0], tt.args[1], tt.args[2])

			assert.Equalf(t, tt.want[0], c, "clipCMY(): C = %v, want %v", c, tt.want[0])
			assert.Equalf(t, tt.want[1], m, "clipCMY(): M = %v, want %v", m, tt.want[1])
			assert.Equalf(t, tt.want[2], y, "clipCMY(): Y = %v, want %v", y, tt.want[2])
		})
	}
}

func TestClip_CMYK(t *testing.T) {
	var tests []struct {
		name string
		args [4]float64
		want [4]float64
	}

	tests = []struct {
		name string
		args [4]float64
		want [4]float64
	}{
		{name: "Under", args: [4]float64{-1.0, -1.0, -1.0, -1.0}, want: [4]float64{CMYKMin, CMYKMin, CMYKMin, CMYKMin}},
		{name: "Normal", args: [4]float64{0.5, 0.5, 0.5, 0.5}, want: [4]float64{0.5, 0.5, 0.5, 0.5}},
		{name: "Over", args: [4]float64{2.0, 2.0, 2.0, 2.0}, want: [4]float64{CMYKMax, CMYKMax, CMYKMax, CMYKMax}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, m, y, k := clipCMYK(tt.args[0], tt.args[1], tt.args[2], tt.args[3])

			assert.Equalf(t, tt.want[0], c, "clipCMYK(): C = %v, want %v", c, tt.want[0])
			assert.Equalf(t, tt.want[1], m, "clipCMYK(): M = %v, want %v", m, tt.want[1])
			assert.Equalf(t, tt.want[2], y, "clipCMYK(): Y = %v, want %v", y, tt.want[2])
			assert.Equalf(t, tt.want[3], k, "clipCMYK(): K = %v, want %v", k, tt.want[3])
		})
	}
}
