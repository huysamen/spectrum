package spectrum

// RGBToCMY converts an RGB color to CMY.
// The RGB color is represented as a 3-tuple of floats in the range [0, 1].
// The CMY color is returned as a 3-tuple of floats in the range [0, 1].
func RGBToCMY(r, g, b float64) (c, m, y float64) {
	return clipCMY(1.0-r, 1.0-g, 1.0-b)
}
