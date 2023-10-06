package spectrum

import "math"

// LabToXYZ converts CIE L*a*b* to CIE XYZ.
// L*a*b* is in the range [0, 100], [-128, 127], [-128, 127].
// CIE XYZ is in the range [0, 0.95047], [0, 1.00000], [0, 1.08883].
func LabToXYZ(l, a, b float64) (x, y, z float64) {
	l, a, b = clipLab(l, a, b)
	x, y, z = LabToXYZWithReference(
		l, a, b,
		ReferenceDaylightRGB02[0],
		ReferenceDaylightRGB02[1],
		ReferenceDaylightRGB02[2],
	)

	return
}

// LabToXYZWithReference converts CIE L*a*b* to CIE XYZ.
// L*a*b* is in the range [0, 100], [-128, 127], [-128, 127].
// CIE XYZ is in the range [0, 0.95047], [0, 1.00000], [0, 1.08883].
// rx, ry, rz are the reference white values.
func LabToXYZWithReference(l, a, b, refX, refY, refZ float64) (x, y, z float64) {
	l, a, b = clipLab(l, a, b)

	y = (l + 16.0) / 116.0
	x = a/500.0 + y
	z = y - b/200.0

	if math.Pow(y, 3.0) > 0.008856 {
		y = math.Pow(y, 3.0) * refY
	} else {
		y = (y - 16.0/116.0) / 7.787 * refY
	}

	if math.Pow(x, 3.0) > 0.008856 {
		x = math.Pow(x, 3.0) * refX
	} else {
		x = (x - 16.0/116.0) / 7.787 * refX
	}

	if math.Pow(z, 3.0) > 0.008856 {
		z = math.Pow(z, 3.0) * refZ
	} else {
		z = (z - 16.0/116.0) / 7.787 * refZ
	}

	x, y, z = clipXYZ(x, y, z)

	return
}
