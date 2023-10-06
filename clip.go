package spectrum

const (
	RGBMin    = float64(0)
	RGBMax    = float64(1.0)
	RGBui8Min = uint8(0)
	RGBui8Max = uint8(255)
	XYZXMin   = float64(0.0)
	XYZXMax   = float64(0.95047)
	XYZYMin   = float64(0.0)
	XYZYMax   = float64(1.00000)
	XYZZMin   = float64(0.0)
	XYZZMax   = float64(1.08883)
	LabLMin   = float64(0.0)
	LabLMax   = float64(100.0)
	LabAMin   = float64(-128.0)
	LabAMax   = float64(127.0)
	LabBMin   = float64(-128.0)
	LabBMax   = float64(127.0)
	HSLMin    = float64(0.0)
	HSLMax    = float64(1.0)
	CMYMin    = float64(0.0)
	CMYMax    = float64(1.0)
	CMYKMin   = float64(0.0)
	CMYKMax   = float64(1.0)
)

func clip[T uint8 | float64](v, min, max T) T {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}

func clipRGB(_r, _g, _b float64) (r, g, b float64) {
	return clip(_r, RGBMin, RGBMax), clip(_g, RGBMin, RGBMax), clip(_b, RGBMin, RGBMax)
}

func clipRGBui8(_r, _g, _b uint8) (r, g, b uint8) {
	return clip(_r, RGBui8Min, RGBui8Max), clip(_g, RGBui8Min, RGBui8Max), clip(_b, RGBui8Min, RGBui8Max)
}

func clipXYZ(_x, _y, _z float64) (x, y, z float64) {
	return clip(_x, XYZXMin, XYZXMax), clip(_y, XYZYMin, XYZYMax), clip(_z, XYZZMin, XYZZMax)
}

func clipLab(_l, _a, _b float64) (l, a, b float64) {
	return clip(_l, LabLMin, LabLMax), clip(_a, LabAMin, LabAMax), clip(_b, LabBMin, LabBMax)
}

func clipHSL(_h, _s, _l float64) (h, s, l float64) {
	return clip(_h, HSLMin, HSLMax), clip(_s, HSLMin, HSLMax), clip(_l, HSLMin, HSLMax)
}

func clipCMY(_c, _m, _y float64) (c, m, y float64) {
	return clip(_c, CMYMin, CMYMax), clip(_m, CMYMin, CMYMax), clip(_y, CMYMin, CMYMax)
}

func clipCMYK(_c, _m, _y, _k float64) (c, m, y, k float64) {
	return clip(_c, CMYKMin, CMYKMax), clip(_m, CMYKMin, CMYKMax), clip(_y, CMYKMin, CMYKMax), clip(_k, CMYKMin, CMYKMax)
}
