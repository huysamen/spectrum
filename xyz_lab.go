package spectrum

import "math"

// XYZToLab converts XYZ to Lab.
// CIE XYZ is in the range [0, 0.95047], [0, 1.00000], [0, 1.08883].
// L*a*b* is in the range [0, 100], [-128, 127], [-128, 127].
func XYZToLab(x, y, z float64) (l, a, b float64) {
	x, y, z = clipXYZ(x, y, z)
	l, a, b = XYZToLabWithObserver(
		x, y, z,
		ReferenceDaylightRGB02[0],
		ReferenceDaylightRGB02[1],
		ReferenceDaylightRGB02[2],
	)

	return
}

// XYZToLabWithObserver converts XYZ to Lab.
// CIE XYZ is in the range [0, 0.95047], [0, 1.00000], [0, 1.08883].
// L*a*b* is in the range [0, 100], [-128, 127], [-128, 127].
// rx, ry, rz are the reference white values.
func XYZToLabWithObserver(x, y, z, refX, refY, refZ float64) (l, a, b float64) {
	x, y, z = clipXYZ(x, y, z)

	x /= refX
	y /= refY
	z /= refZ

	if x > 0.008856 {
		x = math.Pow(x, 1.0/3.0)
	} else {
		x = 7.787*x + 16.0/116.0
	}

	if y > 0.008856 {
		y = math.Pow(y, 1.0/3.0)
	} else {
		y = 7.787*y + 16.0/116.0
	}

	if z > 0.008856 {
		z = math.Pow(z, 1.0/3.0)
	} else {
		z = 7.787*z + 16.0/116.0
	}

	l, a, b = clipLab(116.0*y-16.0, 500.0*(x-y), 200.0*(y-z))

	return
}
