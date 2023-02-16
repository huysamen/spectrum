package spectrum

func CMYToCMYK(c_, m_, y_ float64) (c, m, y, k float64) {
	k = 1.0

	if c_ < k {
		k = c_
	}

	if m_ < k {
		k = m_
	}

	if y_ < k {
		k = y_
	}

	if k == 1.0 {
		return 0.0, 0.0, 0.0, clip(k, CMYKMin, CMYKMax)
	}

	return clipCMYK((c_-k)/(1.0-k), (m_-k)/(1.0-k), (y_-k)/(1.0-k), k)
}
