package spectrum

import "math"

// XYZToRGB converts XYZ to RGB.
// XYZ is in the range [0, 1], RGB is in the range [0, 1].
func XYZToRGB(x, y, z float64) (r, g, b float64) {
	x, y, z = clipXYZ(x, y, z)

	r = 3.240479*x - 1.537150*y - 0.498535*z
	g = -0.969256*x + 1.875992*y + 0.041556*z
	b = 0.055648*x - 0.204043*y + 1.057311*z

	if r > 0.0031308 {
		r = math.Max(0, math.Min(1, 1.055*math.Pow(r, 1.0/2.4)-0.055))
	} else {
		r *= math.Max(0, math.Min(1, 12.92))
	}

	if g > 0.0031308 {
		g = math.Max(0, math.Min(1, 1.055*math.Pow(g, 1.0/2.4)-0.055))
	} else {
		g *= math.Max(0, math.Min(1, 12.92))
	}

	if b > 0.0031308 {
		b = math.Max(0, math.Min(1, 1.055*math.Pow(b, 1.0/2.4)-0.055))
	} else {
		b *= math.Max(0, math.Min(1, 12.92))
	}

	r, g, b = clipRGB(r, g, b)

	return
}
