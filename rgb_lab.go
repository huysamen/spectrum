package spectrum

// RGBToLab converts RGB (in [0,1]) to Lab (L* in [0,100], a* in [-128,127], b* in [-128,127])
func RGBToLab(r, g, b float64) (l, a, b2 float64) {
	x, y, z := RGBToXYZ(r, g, b)
	l, a, b2 = XYZToLab(x, y, z)

	return
}
