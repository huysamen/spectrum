package spectrum

import "math"

// RGBToHSL converts RGB to HSL.
// RGB is in the range [0, 1], HSL is in the range [0, 1].
func RGBToHSL(r, g, b float64) (h, s, l float64) {
	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)
	d := max - min

	l = (max + min) / 2.0

	if d != 0.0 {
		if l < 0.5 {
			s = d / (max + min)
		} else {
			s = d / (2.0 - max - min)
		}

		switch {
		case r == max:
			h = (g - b) / d
		case g == max:
			h = 2.0 + (b-r)/d
		case b == max:
			h = 4.0 + (r-g)/d
		}

		h /= 6.0

		if h < 0.0 {
			h += 1.0
		}
	}

	return clipHSL(h, s, l)
}
