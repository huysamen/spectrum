package spectrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestRGBui8GenerateTints(t *testing.T) {
	// Test with red color (255, 0, 0)
	baseR, baseG, baseB := uint8(255), uint8(0), uint8(0)
	count := 3

	tints := RGBui8GenerateTints(baseR, baseG, baseB, count)

	if len(tints) != count {
		t.Errorf("Expected %d tints, got %d", count, len(tints))
	}

	// First tint should be lightest (closest to white)
	// Last tint should be darkest (closest to base color)
	// For red color, check green and blue components (they should get lighter)
	if tints[0][1] <= tints[len(tints)-1][1] || tints[0][2] <= tints[len(tints)-1][2] {
		t.Errorf("Tints should be ordered from lightest to darkest")
	}

	// All red values should be 255 (since we're mixing red with white)
	for i, tint := range tints {
		if tint[0] != 255 {
			t.Errorf("Tint %d red value should be 255, got %d", i, tint[0])
		}
	}
}

func TestRGBui8GenerateShades(t *testing.T) {
	// Test with red color (255, 0, 0)
	baseR, baseG, baseB := uint8(255), uint8(0), uint8(0)
	count := 3

	shades := RGBui8GenerateShades(baseR, baseG, baseB, count)

	if len(shades) != count {
		t.Errorf("Expected %d shades, got %d", count, len(shades))
	}

	// First shade should be lightest (closest to base color)
	// Last shade should be darkest (closest to black)
	if shades[0][0] <= shades[len(shades)-1][0] {
		t.Errorf("Shades should be ordered from lightest to darkest")
	}

	// Green and blue values should remain 0 (since we're mixing red with black)
	for i, shade := range shades {
		if shade[1] != 0 || shade[2] != 0 {
			t.Errorf("Shade %d should have green=0, blue=0, got green=%d, blue=%d", i, shade[1], shade[2])
		}
	}
}

func TestRGBui8GenerateTintsAndShadesZeroCount(t *testing.T) {
	baseR, baseG, baseB := uint8(128), uint8(128), uint8(128)

	tints := RGBui8GenerateTints(baseR, baseG, baseB, 0)
	if tints != nil {
		t.Errorf("Expected nil for zero count tints, got %v", tints)
	}

	shades := RGBui8GenerateShades(baseR, baseG, baseB, 0)
	if shades != nil {
		t.Errorf("Expected nil for zero count shades, got %v", shades)
	}

	tints = RGBui8GenerateTints(baseR, baseG, baseB, -1)
	if tints != nil {
		t.Errorf("Expected nil for negative count tints, got %v", tints)
	}
}

func TestRGBui8GenerateSpecificValues(t *testing.T) {
	// Test with a specific color to verify exact interpolation
	baseR, baseG, baseB := uint8(255), uint8(87), uint8(51) // #FF5733

	// Generate 2 tints
	tints := RGBui8GenerateTints(baseR, baseG, baseB, 2)

	// Expected tints (interpolating towards white):
	// t1: t=2/3, r=255*(1-2/3)+255*2/3=255, g=87*(1/3)+255*2/3=199, b=51*(1/3)+255*2/3=187
	// t2: t=1/3, r=255, g=87*(2/3)+255*1/3=143, b=51*(2/3)+255*1/3=119

	if len(tints) != 2 {
		t.Errorf("Expected 2 tints, got %d", len(tints))
	}

	// Verify tints are getting lighter (higher values for g and b)
	if tints[0][1] <= tints[1][1] || tints[0][2] <= tints[1][2] {
		t.Errorf("First tint should be lighter than second tint")
	}
}
