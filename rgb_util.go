package spectrum

import "math"

// RGBToRGBui8 converts RGB to RGBui8.
// RGB is in the range [0, 1], RGBui8 is in the range [0, 255].
func RGBToRGBui8(_r, _g, _b float64) (r, g, b uint8) {
	_r, _g, _b = clipRGB(_r, _g, _b)
	r, g, b = clipRGBui8(uint8(math.Round(_r*255.0)), uint8(math.Round(_g*255.0)), uint8(math.Round(_b*255.0)))

	return
}

// RGBui8ToRGB converts RGBui8 to RGB.
// RGBui8 is in the range [0, 255], RGB is in the range [0, 1].
func RGBui8ToRGB(_r, _g, _b uint8) (r, g, b float64) {
	_r, _g, _b = clipRGBui8(_r, _g, _b)
	r, g, b = clipRGB(float64(_r)/255.0, float64(_g)/255.0, float64(_b)/255.0)

	return
}

// RGBToRGBui32 converts RGB to RGBui32.
// RGB is in the range [0, 1], RGBui32 is integer representation of RGB.
func RGBToRGBui32(r, g, b float64) (rgb uint32) {
	r, g, b = clipRGB(r, g, b)
	rgb = uint32(math.Round(r*255.0))<<16 | uint32(math.Round(g*255.0))<<8 | uint32(math.Round(b*255.0))

	return
}

// RGBui32ToRGB converts RGBui32 to RGB.
// RGBui32 is integer representation of RGB, RGB is in the range [0, 1].
func RGBui32ToRGB(rgb uint32) (r, g, b float64) {
	r, g, b = clipRGB(float64((rgb>>16)&0xff)/255.0, float64((rgb>>8)&0xff)/255.0, float64(rgb&0xff)/255.0)

	return
}
