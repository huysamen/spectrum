package spectrum

// RGBToRGBui8 converts RGB to RGBui8.
// RGB is in the range [0, 1], RGBui8 is in the range [0, 255].
func RGBToRGBui8(r, g, b float64) (uint8, uint8, uint8) {
	return clipRGBui8(uint8(r*255.0), uint8(g*255.0), uint8(b*255.0))
}

// RGBui8ToRGB converts RGBui8 to RGB.
// RGBui8 is in the range [0, 255], RGB is in the range [0, 1].
func RGBui8ToRGB(r, g, b uint8) (float64, float64, float64) {
	return clipRGB(float64(r)/255.0, float64(g)/255.0, float64(b)/255.0)
}

// RGBToRGBui32 converts RGB to RGBui32.
// RGB is in the range [0, 1], RGBui32 is integer representation of RGB.
func RGBToRGBui32(r, g, b float64) uint32 {
	r, g, b = clipRGB(r, g, b)

	return uint32(r*255.0)<<16 | uint32(g*255.0)<<8 | uint32(b*255.0)
}

// RGBui32ToRGB converts RGBui32 to RGB.
// RGBui32 is integer representation of RGB, RGB is in the range [0, 1].
func RGBui32ToRGB(rgb uint32) (float64, float64, float64) {
	return clipRGB(
		float64((rgb>>16)&0xff)/255.0,
		float64((rgb>>8)&0xff)/255.0,
		float64(rgb&0xff)/255.0,
	)
}
