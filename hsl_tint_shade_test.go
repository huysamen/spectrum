package spectrum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTints(t *testing.T) {
	tests := []struct {
		name     string
		h, s, l  float64
		number   int
		expected int
	}{
		{
			name:     "Generate 3 tints from medium blue",
			h:        0.6667, // Blue hue
			s:        1.0,    // Full saturation
			l:        0.5,    // Medium lightness
			number:   3,
			expected: 3, // 3 tints (no base color)
		},
		{
			name:     "Generate 1 tint",
			h:        0.0, // Red hue
			s:        1.0, // Full saturation
			l:        0.3, // Dark
			number:   1,
			expected: 1,
		},
		{
			name:     "Generate 0 tints",
			h:        0.0,
			s:        1.0,
			l:        0.5,
			number:   0,
			expected: 0,
		},
		{
			name:     "Generate negative number of tints",
			h:        0.0,
			s:        1.0,
			l:        0.5,
			number:   -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateTints(tt.h, tt.s, tt.l, tt.number)
			assert.Equal(t, tt.expected, len(result))

			if len(result) > 0 {
				// All tints should have lightness greater than base color
				for i, tint := range result {
					assert.True(t, tint[2] > tt.l, "Tint %d lightness should be greater than base lightness", i)
					assert.True(t, tint[2] <= 1.0, "Lightness should not exceed 1.0")
					assert.Equal(t, tt.h, tint[0], "Hue should remain constant")
					assert.Equal(t, tt.s, tint[1], "Saturation should remain constant")
				}

				// Lightness should increase for tints
				for i := 1; i < len(result); i++ {
					assert.True(t, result[i][2] >= result[i-1][2], "Lightness should increase for tints")
				}
			}
		})
	}
}

func TestGenerateShades(t *testing.T) {
	tests := []struct {
		name     string
		h, s, l  float64
		number   int
		expected int
	}{
		{
			name:     "Generate 3 shades from medium blue",
			h:        0.6667, // Blue hue
			s:        1.0,    // Full saturation
			l:        0.5,    // Medium lightness
			number:   3,
			expected: 3, // 3 shades (no base color)
		},
		{
			name:     "Generate 1 shade",
			h:        0.0, // Red hue
			s:        1.0, // Full saturation
			l:        0.7, // Light
			number:   1,
			expected: 1,
		},
		{
			name:     "Generate 0 shades",
			h:        0.0,
			s:        1.0,
			l:        0.5,
			number:   0,
			expected: 0,
		},
		{
			name:     "Generate negative number of shades",
			h:        0.0,
			s:        1.0,
			l:        0.5,
			number:   -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateShades(tt.h, tt.s, tt.l, tt.number)
			assert.Equal(t, tt.expected, len(result))

			if len(result) > 0 {
				// All shades should have lightness less than base color
				for i, shade := range result {
					assert.True(t, shade[2] < tt.l, "Shade %d lightness should be less than base lightness", i)
					assert.True(t, shade[2] >= 0.0, "Lightness should not go below 0.0")
					assert.Equal(t, tt.h, shade[0], "Hue should remain constant")
					assert.Equal(t, tt.s, shade[1], "Saturation should remain constant")
				}

				// Lightness should decrease for shades
				for i := 1; i < len(result); i++ {
					assert.True(t, result[i][2] <= result[i-1][2], "Lightness should decrease for shades")
				}
			}
		})
	}
}

func TestGenerateSpecificTintsAndShades(t *testing.T) {
	t.Run("Generate specific tints with exact lightness values", func(t *testing.T) {
		// Base color: Red with 50% lightness
		h, s, l := 0.6194, 1.0, 0.61
		tints := GenerateTints(h, s, l, 5)

		assert.Equal(t, 5, len(tints))

		expected := []float64{0.675, 0.74, 0.805, 0.87, 0.935}

		for i, expectedL := range expected {
			assert.InDelta(t, expectedL, tints[i][2], 0.0001, "Tint %d should have lightness %.3f", i, expectedL)
			assert.Equal(t, h, tints[i][0], "Hue should remain constant")
			assert.Equal(t, s, tints[i][1], "Saturation should remain constant")

			rr, gg, bb := HSLToRGB(tints[i][0], tints[i][1], tints[i][2])
			ri, gi, bi := RGBToRGBui8(rr, gg, bb)
			hex := RGBui8ToHexString(ri, gi, bi)

			fmt.Printf("Tint %d: H=%.4f, S=%.4f, L=%.4f, RGB=(%.2f, %.2f, %.2f), Hex=%s\n",
				i, tints[i][0], tints[i][1], tints[i][2],
				rr, gg, bb, hex)
		}
	})

	t.Run("Generate specific shades with exact lightness values", func(t *testing.T) {
		// Base color: Red with 50% lightness
		h, s, l := 0.6194, 1.0, 0.61
		shades := GenerateShades(h, s, l, 5)

		assert.Equal(t, 5, len(shades))

		expected := []float64{0.5083, 0.4067, 0.305, 0.2033, 0.10167}

		for i, expectedL := range expected {
			assert.InDelta(t, expectedL, shades[i][2], 0.0001, "Shade %d should have lightness %.3f", i, expectedL)
			assert.Equal(t, h, shades[i][0], "Hue should remain constant")
			assert.Equal(t, s, shades[i][1], "Saturation should remain constant")

			rr, gg, bb := HSLToRGB(shades[i][0], shades[i][1], shades[i][2])
			ri, gi, bi := RGBToRGBui8(rr, gg, bb)
			hex := RGBui8ToHexString(ri, gi, bi)

			fmt.Printf("Shade %d: H=%.4f, S=%.4f, L=%.4f, RGB=(%.2f, %.2f, %.2f), Hex=%s\n",
				i, shades[i][0], shades[i][1], shades[i][2],
				rr, gg, bb, hex)
		}
	})
}

func TestGenerateTintsAndShadesEdgeCases(t *testing.T) {
	t.Run("Tints from very light color", func(t *testing.T) {
		tints := GenerateTints(0.0, 1.0, 0.9, 2)
		assert.Equal(t, 2, len(tints))
		// With number+1 divisor, lightness = 0.9 + (1.0-0.9)/3 * 2 = 0.9 + 0.1/3 * 2 = 0.9 + 0.0667 ≈ 0.9667
		assert.InDelta(t, 0.9667, tints[1][2], 0.001, "Should not reach maximum lightness")
	})

	t.Run("Shades from very dark color", func(t *testing.T) {
		shades := GenerateShades(0.0, 1.0, 0.1, 2)
		assert.Equal(t, 2, len(shades))
		// With number+1 divisor, lightness = 0.1 - 0.1/3 * 2 = 0.1 - 0.0667 ≈ 0.0333
		assert.InDelta(t, 0.0333, shades[1][2], 0.001, "Should not reach minimum lightness")
	})

	t.Run("Input values are clipped", func(t *testing.T) {
		// Test with out-of-range input values - base lightness is clipped to 1.0,
		// so tints will be very close to but not at maximum lightness
		tints := GenerateTints(-0.5, 1.5, 1.5, 2)
		assert.Equal(t, 2, len(tints))
		assert.Equal(t, 1.0, tints[0][2]) // All tints will be at max lightness since base is already at 1.0
		assert.Equal(t, 1.0, tints[1][2])
		// Check that hue and saturation are clipped correctly
		assert.Equal(t, 0.0, tints[0][0]) // Hue clipped to 0.0
		assert.Equal(t, 1.0, tints[0][1]) // Saturation clipped to 1.0
	})
}
