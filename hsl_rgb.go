package spectrum

// HSLToRGB converts HSL to RGB.
// HSL is in the range [0, 1], RGB is in the range [0, 1].
func HSLToRGB(h, s, l float64) (r, g, b float64) {
	if s == 0 {
		return l, l, l
	}

	var q, p float64

	if l < 0.5 {
		q = l * (1.0 + s)
	} else {
		q = (l + s) - l*s
	}

	p = 2.0*l - q

	r = hueToRGB(p, q, h+1.0/3.0)
	g = hueToRGB(p, q, h)
	b = hueToRGB(p, q, h-1.0/3.0)

	return clipRGB(r, g, b)
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0.0 {
		t += 1.0
	}

	if t > 1.0 {
		t -= 1.0
	}

	if t*6.0 < 1.0 {
		return p + (q-p)*6.0*t
	}

	if t*2.0 < 1.0 {
		return q
	}

	if t*3.0 < 2.0 {
		return p + (q-p)*((2.0/3.0)-t)*6.0
	}

	return p
}
