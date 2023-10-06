package spectrum

// CMYKToCMY converts a CMYK color to CMY.
// The CMYK color is returned as a 4-tuple of floats in the range [0, 1].
// The CMY color is represented as a 3-tuple of floats in the range [0, 1].
func CMYKToCMY(_c, _m, _y, _k float64) (c, m, y float64) {
	_c, _m, _y = clipCMY(_c, _m, _y)
	c, m, y = clipCMY((1.0-_k)*_c+_k, (1.0-_k)*_m+_k, (1.0-_k)*_y+_k)

	return
}
