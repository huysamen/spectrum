package spectrum

// CMYToRGB converts a CMY color to RGB.
// The CMY color is represented as a 3-tuple of floats in the range [0, 1].
// The RGB color is returned as a 3-tuple of floats in the range [0, 1].
func CMYToRGB(c, m, y float64) (r, g, b float64) {
	c, m, y = clipCMY(c, m, y)

	return clipRGB(1.0-c, 1.0-m, 1.0-y)
}
