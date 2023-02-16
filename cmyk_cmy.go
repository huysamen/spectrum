package spectrum

func CMYKToCMY(c_, m_, y_, k_ float64) (c, m, y float64) {
	return clipCMY(
		(1.0-k_)*c_+k_,
		(1.0-k_)*m_+k_,
		(1.0-k_)*y_+k_,
	)
}
