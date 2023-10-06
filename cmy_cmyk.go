package spectrum

import "math"

// CMYToCMYK converts a CMY color to CMYK.
// The CMY color is represented as a 3-tuple of floats in the range [0, 1].
// The CMYK color is returned as a 4-tuple of floats in the range [0, 1].
func CMYToCMYK(_c, _m, _y float64) (c, m, y, k float64) {
	_c, _m, _y = clipCMY(_c, _m, _y)
	k = math.Min(math.Min(_c, _m), _y)

	if k == 1.0 {
		return 0.0, 0.0, 0.0, 1.0
	}

	c = (_c - k) / (1.0 - k)
	m = (_m - k) / (1.0 - k)
	y = (_y - k) / (1.0 - k)

	return clipCMYK(c, m, y, k)
}
