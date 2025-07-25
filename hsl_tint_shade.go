package spectrum

// HSLGenerateTints generates a slice of tints for the given HSL color.
// Tints are created by increasing the lightness value towards white (1.0).
// The base color HSL values should be in the range [0, 1].
// The number parameter determines how many tints to generate and the step size.
// Returns a slice of HSL values, excluding the base color.
func HSLGenerateTints(h, s, l float64, number int) [][3]float64 {
	if number <= 0 {
		return [][3]float64{}
	}

	h, s, l = clipHSL(h, s, l)
	tints := make([][3]float64, number)

	const (
		maxLightnessIncrease = 0.95 // How much to move towards white (increased from 0.8)
		saturationReduction  = 0.8  // Reduce saturation as we get lighter (increased from 0.6)
	)

	for i := 1; i <= number; i++ {
		// Linear progression
		step := float64(i) / float64(number+1) // +1 prevents reaching pure white

		// Keep hue constant, reduce saturation, increase lightness
		newH := h
		newS := s * (1.0 - saturationReduction*step)
		newL := l + (1.0-l)*maxLightnessIncrease*step

		tints[i-1] = [3]float64{newH, newS, newL}
	}

	return tints
}

// HSLGenerateShades generates a slice of shades for the given HSL color.
// Shades are created by decreasing the lightness value towards black (0.0).
// The base color HSL values should be in the range [0, 1].
// The number parameter determines how many shades to generate and the step size.
// Returns a slice of HSL values, excluding the base color.
func HSLGenerateShades(h, s, l float64, number int) [][3]float64 {
	if number <= 0 {
		return [][3]float64{}
	}

	h, s, l = clipHSL(h, s, l)
	shades := make([][3]float64, number)

	const (
		maxLightnessReduction = 1.0 // Allow going down to 5% of original (increased from 0.9)
		saturationBoost       = 0.1 // Slight saturation increase for darker shades (increased from 0.1)
	)

	for i := 1; i <= number; i++ {
		// Linear progression
		step := float64(i) / float64(number+1) // +1 prevents reaching pure black

		// Keep hue constant, slightly boost saturation, reduce lightness
		newH := h
		newS := clip(s*(1.0+saturationBoost*step), 0.0, 1.0)
		newL := l * (1.0 - maxLightnessReduction*step)

		shades[i-1] = [3]float64{newH, newS, newL}
	}

	return shades
}
