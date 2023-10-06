package spectrum

import "math"

// RGBToXYZ converts RGB to XYZ.
// RGB is in the range [0, 1], XYZ is in the range [0, 1].
func RGBToXYZ(r, g, b float64) (x, y, z float64) {
	r, g, b = clipRGB(r, g, b)

	if r > 0.04045 {
		r = math.Pow((r+0.055)/1.055, 2.4)
	} else {
		r /= 12.92
	}

	if g > 0.04045 {
		g = math.Pow((g+0.055)/1.055, 2.4)
	} else {
		g /= 12.92
	}

	if b > 0.04045 {
		b = math.Pow((b+0.055)/1.055, 2.4)
	} else {
		b /= 12.92
	}

	x, y, z = clipXYZ(
		0.412453*r+0.357580*g+0.180423*b,
		0.212671*r+0.715160*g+0.072169*b,
		0.019334*r+0.119193*g+0.950227*b,
	)

	return
}
