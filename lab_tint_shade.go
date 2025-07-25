package spectrum

// LABGenerateTints generates a slice of tints for the given LAB color.
// Tints are created by increasing the L* value towards white (100.0).
// The base color LAB values should be in their respective ranges:
// L*: [0, 100], a*: [-128, 127], b*: [-128, 127].
// The number parameter determines how many tints to generate.
// Returns a slice of LAB values, excluding the base color.
func LABGenerateTints(l, a, b float64, number int) [][3]float64 {
	if number <= 0 {
		return [][3]float64{}
	}

	l, a, b = clipLab(l, a, b)
	tints := make([][3]float64, number)

	const (
		maxLightnessIncrease = 0.95 // How much to move towards white (L* = 100)
		chromaReduction      = 0.8  // Reduce chroma (a*, b*) as we get lighter
	)

	for i := 1; i <= number; i++ {
		// Linear progression
		step := float64(i) / float64(number+1) // +1 prevents reaching pure white

		// Increase L* towards 100, reduce chroma (a*, b*) towards 0
		newL := l + (100.0-l)*maxLightnessIncrease*step
		newA := a * (1.0 - chromaReduction*step)
		newB := b * (1.0 - chromaReduction*step)

		tints[i-1] = [3]float64{newL, newA, newB}
	}

	return tints
}

// LABGenerateShades generates a slice of shades for the given LAB color.
// Shades are created by decreasing the L* value towards black (0.0).
// The base color LAB values should be in their respective ranges:
// L*: [0, 100], a*: [-128, 127], b*: [-128, 127].
// The number parameter determines how many shades to generate.
// Returns a slice of LAB values, excluding the base color.
func LABGenerateShades(l, a, b float64, number int) [][3]float64 {
	if number <= 0 {
		return [][3]float64{}
	}

	l, a, b = clipLab(l, a, b)
	shades := make([][3]float64, number)

	const (
		maxLightnessReduction = 0.95 // Allow going down to 5% of original L*
		chromaBoost           = 0.2  // Slight chroma increase for darker shades
	)

	for i := 1; i <= number; i++ {
		// Linear progression
		step := float64(i) / float64(number+1) // +1 prevents reaching pure black

		// Decrease L* towards 0, slightly boost chroma (a*, b*)
		newL := l * (1.0 - maxLightnessReduction*step)
		newA := clip(a*(1.0+chromaBoost*step), LabAMin, LabAMax)
		newB := clip(b*(1.0+chromaBoost*step), LabBMin, LabBMax)

		shades[i-1] = [3]float64{newL, newA, newB}
	}

	return shades
}

// LABGenerateTintsRGB is a convenience function that generates tints in LAB space
// but returns RGB values for easier use in applications.
func LABGenerateTintsRGB(r, g, b float64, number int) [][3]float64 {
	// Convert RGB to LAB
	x, y, z := RGBToXYZ(r, g, b)
	l, a, labB := XYZToLab(x, y, z)

	// Generate tints in LAB space
	labTints := LABGenerateTints(l, a, labB, number)

	// Convert back to RGB
	rgbTints := make([][3]float64, len(labTints))
	for i, tint := range labTints {
		x, y, z := LabToXYZ(tint[0], tint[1], tint[2])
		r, g, b := XYZToRGB(x, y, z)
		rgbTints[i] = [3]float64{r, g, b}
	}

	return rgbTints
}

// LABGenerateShadesRGB is a convenience function that generates shades in LAB space
// but returns RGB values for easier use in applications.
func LABGenerateShadesRGB(r, g, b float64, number int) [][3]float64 {
	// Convert RGB to LAB
	x, y, z := RGBToXYZ(r, g, b)
	l, a, labB := XYZToLab(x, y, z)

	// Generate shades in LAB space
	labShades := LABGenerateShades(l, a, labB, number)

	// Convert back to RGB
	rgbShades := make([][3]float64, len(labShades))
	for i, shade := range labShades {
		x, y, z := LabToXYZ(shade[0], shade[1], shade[2])
		r, g, b := XYZToRGB(x, y, z)
		rgbShades[i] = [3]float64{r, g, b}
	}

	return rgbShades
}
