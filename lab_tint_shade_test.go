package spectrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLABGenerateTints(t *testing.T) {
	tests := []struct {
		name     string
		l, a, b  float64
		number   int
		expected int
	}{
		{
			name:     "Generate 3 tints from medium blue LAB",
			l:        50.0,  // Medium lightness
			a:        20.0,  // Green-Red axis
			b:        -40.0, // Blue-Yellow axis (negative = blue)
			number:   3,
			expected: 3,
		},
		{
			name:     "Generate 1 tint",
			l:        30.0,
			a:        50.0,
			b:        20.0,
			number:   1,
			expected: 1,
		},
		{
			name:     "Generate 0 tints",
			l:        50.0,
			a:        0.0,
			b:        0.0,
			number:   0,
			expected: 0,
		},
		{
			name:     "Generate negative number of tints",
			l:        50.0,
			a:        0.0,
			b:        0.0,
			number:   -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LABGenerateTints(tt.l, tt.a, tt.b, tt.number)
			assert.Equal(t, tt.expected, len(result))

			if len(result) > 0 {
				// All tints should have lightness greater than base color
				for i, tint := range result {
					assert.True(t, tint[0] > tt.l, "Tint %d lightness should be greater than base lightness", i)
					assert.True(t, tint[0] <= 100.0, "Lightness should not exceed 100.0")
				}

				// Lightness should increase for tints
				for i := 1; i < len(result); i++ {
					assert.True(t, result[i][0] >= result[i-1][0], "Lightness should increase for tints")
				}

				// Chroma (a*, b*) should decrease as we approach white
				for i, tint := range result {
					chromaBase := tt.a*tt.a + tt.b*tt.b
					chromaTint := tint[1]*tint[1] + tint[2]*tint[2]
					assert.True(t, chromaTint <= chromaBase, "Tint %d should have lower or equal chroma", i)
				}
			}
		})
	}
}

func TestLABGenerateShades(t *testing.T) {
	tests := []struct {
		name     string
		l, a, b  float64
		number   int
		expected int
	}{
		{
			name:     "Generate 3 shades from medium blue LAB",
			l:        50.0,  // Medium lightness
			a:        20.0,  // Green-Red axis
			b:        -40.0, // Blue-Yellow axis (negative = blue)
			number:   3,
			expected: 3,
		},
		{
			name:     "Generate 1 shade",
			l:        70.0,
			a:        30.0,
			b:        10.0,
			number:   1,
			expected: 1,
		},
		{
			name:     "Generate 0 shades",
			l:        50.0,
			a:        0.0,
			b:        0.0,
			number:   0,
			expected: 0,
		},
		{
			name:     "Generate negative number of shades",
			l:        50.0,
			a:        0.0,
			b:        0.0,
			number:   -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LABGenerateShades(tt.l, tt.a, tt.b, tt.number)
			assert.Equal(t, tt.expected, len(result))

			if len(result) > 0 {
				// All shades should have lightness less than base color
				for i, shade := range result {
					assert.True(t, shade[0] < tt.l, "Shade %d lightness should be less than base lightness", i)
					assert.True(t, shade[0] >= 0.0, "Lightness should not go below 0.0")
				}

				// Lightness should decrease for shades
				for i := 1; i < len(result); i++ {
					assert.True(t, result[i][0] <= result[i-1][0], "Lightness should decrease for shades")
				}
			}
		})
	}
}

func TestLABGenerateTintsRGB(t *testing.T) {
	t.Run("Generate RGB tints using LAB space", func(t *testing.T) {
		// Test with a red color
		r, g, b := 0.8, 0.2, 0.2
		tints := LABGenerateTintsRGB(r, g, b, 3)

		assert.Equal(t, 3, len(tints))

		// All tints should be valid RGB values
		for i, tint := range tints {
			assert.True(t, tint[0] >= 0.0 && tint[0] <= 1.0, "Tint %d R should be in [0,1]", i)
			assert.True(t, tint[1] >= 0.0 && tint[1] <= 1.0, "Tint %d G should be in [0,1]", i)
			assert.True(t, tint[2] >= 0.0 && tint[2] <= 1.0, "Tint %d B should be in [0,1]", i)

			// Tints should generally be lighter (higher values)
			assert.True(t, tint[0] >= r || tint[1] >= g || tint[2] >= b, "Tint %d should be lighter", i)
		}
	})
}

func TestLABGenerateShadesRGB(t *testing.T) {
	t.Run("Generate RGB shades using LAB space", func(t *testing.T) {
		// Test with a bright green color
		r, g, b := 0.2, 0.8, 0.2
		shades := LABGenerateShadesRGB(r, g, b, 3)

		assert.Equal(t, 3, len(shades))

		// All shades should be valid RGB values
		for i, shade := range shades {
			assert.True(t, shade[0] >= 0.0 && shade[0] <= 1.0, "Shade %d R should be in [0,1]", i)
			assert.True(t, shade[1] >= 0.0 && shade[1] <= 1.0, "Shade %d G should be in [0,1]", i)
			assert.True(t, shade[2] >= 0.0 && shade[2] <= 1.0, "Shade %d B should be in [0,1]", i)

			// Shades should generally be darker (lower values)
			assert.True(t, shade[0] <= r || shade[1] <= g || shade[2] <= b, "Shade %d should be darker", i)
		}
	})
}

func TestLABEdgeCases(t *testing.T) {
	t.Run("Tints from very light LAB color", func(t *testing.T) {
		tints := LABGenerateTints(90.0, 5.0, 5.0, 2)
		assert.Equal(t, 2, len(tints))
		// Should not exceed L* = 100
		assert.True(t, tints[1][0] <= 100.0, "Should not exceed maximum lightness")
	})

	t.Run("Shades from very dark LAB color", func(t *testing.T) {
		shades := LABGenerateShades(10.0, 5.0, 5.0, 2)
		assert.Equal(t, 2, len(shades))
		// Should not go below L* = 0
		assert.True(t, shades[1][0] >= 0.0, "Should not go below minimum lightness")
	})

	t.Run("Input values are clipped", func(t *testing.T) {
		// Test with out-of-range input values
		tints := LABGenerateTints(150.0, 200.0, -200.0, 2)
		assert.Equal(t, 2, len(tints))
		// Values should be clipped to valid LAB ranges
		assert.True(t, tints[0][0] <= 100.0, "L* should be clipped to max")
	})
}
