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

// RGBui8Interpolate performs linear interpolation between two RGBui8 colors.
func RGBui8Interpolate(r1, g1, b1, r2, g2, b2 uint8, t float64) (r, g, b uint8) {
	r1, g1, b1 = clipRGBui8(r1, g1, b1)
	r2, g2, b2 = clipRGBui8(r2, g2, b2)

	t = math.Max(0, math.Min(1, t))

	r = uint8(float64(r1)*(1-t) + float64(r2)*t)
	g = uint8(float64(g1)*(1-t) + float64(g2)*t)
	b = uint8(float64(b1)*(1-t) + float64(b2)*t)

	return
}

// RGBui8GenerateTints generates a specified number of tints by interpolating from the base color to white.
// Returns a slice of RGB uint8 color tuples, ordered from lightest to darkest.
func RGBui8GenerateTints(baseR, baseG, baseB uint8, count int) [][3]uint8 {
	if count <= 0 {
		return nil
	}

	tints := make([][3]uint8, count)

	for i := 0; i < count; i++ {
		// Calculate interpolation factor: starts high (closer to white) and decreases
		t := float64(count-i) / float64(count+1)
		r, g, b := RGBui8Interpolate(baseR, baseG, baseB, 255, 255, 255, t)
		tints[i] = [3]uint8{r, g, b}
	}

	return tints
}

// RGBui8GenerateShades generates a specified number of shades by interpolating from the base color to black.
// Returns a slice of RGB uint8 color tuples, ordered from lightest to darkest.
func RGBui8GenerateShades(baseR, baseG, baseB uint8, count int) [][3]uint8 {
	if count <= 0 {
		return nil
	}

	shades := make([][3]uint8, count)

	for i := 0; i < count; i++ {
		// Calculate interpolation factor: starts low (closer to base) and increases
		t := float64(i+1) / float64(count+1)
		r, g, b := RGBui8Interpolate(baseR, baseG, baseB, 0, 0, 0, t)
		shades[i] = [3]uint8{r, g, b}
	}

	return shades
}
