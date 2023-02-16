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

type colourVal interface {
	uint8 | float64
}

func clip[T colourVal](v, min, max T) T {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}

func clipRGB(r, g, b float64) (float64, float64, float64) {
	return clip(r, RGBMin, RGBMax),
		clip(g, RGBMin, RGBMax),
		clip(b, RGBMin, RGBMax)
}

func clipRGBui8(r, g, b uint8) (uint8, uint8, uint8) {
	return clip(r, RGBui8Min, RGBui8Max),
		clip(g, RGBui8Min, RGBui8Max),
		clip(b, RGBui8Min, RGBui8Max)
}

func clipXYZ(x, y, z float64) (float64, float64, float64) {
	return clip(x, XYZXMin, XYZXMax),
		clip(y, XYZYMin, XYZYMax),
		clip(z, XYZZMin, XYZZMax)
}

func clipLab(l, a, b float64) (float64, float64, float64) {
	return clip(l, LabLMin, LabLMax),
		clip(a, LabAMin, LabAMax),
		clip(b, LabBMin, LabBMax)
}

func clipHSL(h, s, l float64) (float64, float64, float64) {
	return clip(h, HSLMin, HSLMax),
		clip(s, HSLMin, HSLMax),
		clip(l, HSLMin, HSLMax)
}

func clipCMY(c, m, y float64) (float64, float64, float64) {
	return clip(c, CMYMin, CMYMax),
		clip(m, CMYMin, CMYMax),
		clip(y, CMYMin, CMYMax)
}

func clipCMYK(c, m, y, k float64) (float64, float64, float64, float64) {
	return clip(c, CMYKMin, CMYKMax),
		clip(m, CMYKMin, CMYKMax),
		clip(y, CMYKMin, CMYKMax),
		clip(k, CMYKMin, CMYKMax)
}
