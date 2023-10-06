package spectrum

import "math"

func DeltaE(l1, a1, b1, l2, a2, b2 float64) (deltaE float64) {
	l1, a1, b1 = clipLab(l1, a1, b1)
	l2, a2, b2 = clipLab(l2, a2, b2)

	deltaE = math.Sqrt(math.Pow(l1-l2, 2.0) + math.Pow(a1-a2, 2.0) + math.Pow(b1-b2, 2.0))

	return
}

func DeltaE94(l1, a1, b1, l2, a2, b2 float64) (deltaE float64) {
	l1, a1, b1 = clipLab(l1, a1, b1)
	l2, a2, b2 = clipLab(l2, a2, b2)

	deltaE = deltaE94WithApplication(l1, a1, b1, l2, a2, b2, 1.0, 0.045, 0.015)

	return
}

func DeltaE94ForGraphicArts(l1, a1, b1, l2, a2, b2 float64) (deltaE float64) {
	l1, a1, b1 = clipLab(l1, a1, b1)
	l2, a2, b2 = clipLab(l2, a2, b2)

	deltaE = deltaE94WithApplication(l1, a1, b1, l2, a2, b2, 1.0, 0.045, 0.015)

	return
}

func DeltaE94ForTextiles(l1, a1, b1, l2, a2, b2 float64) (deltaE float64) {
	l1, a1, b1 = clipLab(l1, a1, b1)
	l2, a2, b2 = clipLab(l2, a2, b2)

	deltaE = deltaE94WithApplication(l1, a1, b1, l2, a2, b2, 2.0, 0.048, 0.014)

	return
}

func deltaE94WithApplication(l1, a1, b1, l2, a2, b2, kL, K1, K2 float64) (deltaE float64) {
	l1, a1, b1 = clipLab(l1, a1, b1)
	l2, a2, b2 = clipLab(l2, a2, b2)

	kC, kH := 1.0, 1.0

	deltaA := a1 - a2
	deltaB := b1 - b2

	C1 := math.Sqrt(a1*a1 + b1*b1)
	C2 := math.Sqrt(a2*a2 + b2*b2)
	deltaC := C1 - C2

	deltaH := math.Sqrt(math.Max(0, deltaA*deltaA+deltaB*deltaB-deltaC*deltaC))

	SL := 1.0
	SC := 1.0 + K1*C1
	SH := 1.0 + K2*C1

	deltaLkLsl := (l1 - l2) / (kL * SL)
	deltaCkSC := deltaC / (SC * kC)
	deltaHkSH := deltaH / (SH * kH)

	deltaE = math.Sqrt(deltaLkLsl*deltaLkLsl + deltaCkSC*deltaCkSC + deltaHkSH*deltaHkSH)

	return
}

// DeltaE2000 is the CIEDE2000 color difference formula.
// See http://www.ece.rochester.edu/~gsharma/ciede2000/ciede2000noteCRNA.pdf
func DeltaE2000(l1, a1, b1, l2, a2, b2 float64) (deltaE float64) {
	l1, a1, b1 = clipLab(l1, a1, b1)
	l2, a2, b2 = clipLab(l2, a2, b2)

	deltaE = DeltaE2000WithWeighingFactors(l1, a1, b1, l2, a2, b2, 1.0, 1.0, 1.0)

	return
}

func DeltaE2000WithWeighingFactors(l1, a1, b1, l2, a2, b2, kL, kC, kH float64) (deltaE float64) {
	l1, a1, b1 = clipLab(l1, a1, b1)
	l2, a2, b2 = clipLab(l2, a2, b2)

	// 1 - Calculate C'i and h'i
	//

	// (2, 3)
	Cb := (math.Sqrt(a1*a1+b1*b1) + math.Sqrt(a2*a2+b2*b2)) / 2.0
	C7 := math.Pow(Cb, 7.0)

	// (4)
	G := 0.5 * (1.0 - math.Sqrt(C7/(C7+6103515625.0)))

	// (5)
	a1_ := (1.0 + G) * a1
	a2_ := (1.0 + G) * a2

	// (6)
	C1_ := math.Sqrt(a1_*a1_ + b1*b1)
	C2_ := math.Sqrt(a2_*a2_ + b2*b2)

	// (7)
	var h1_, h2_ float64

	if b1 == 0.0 && a1_ == 0.0 {
		h1_ = 0.0
	} else {
		h1_ = math.Atan2(rad(b1), rad(a1_))
	}

	if h1_ < 0.0 {
		h1_ += 2.0 * math.Pi
	}

	h1_ *= 180.0 / math.Pi

	if b2 != 0.0 || a2_ != 0.0 {
		h2_ = math.Atan2(rad(b2), rad(a2_))
	}

	if h2_ < 0.0 {
		h2_ += 2.0 * math.Pi
	}

	h2_ *= 180.0 / math.Pi

	// 2 - Calculate dL', dC', dH'
	//

	// (8)
	dL := l2 - l1

	// (9)
	dC := C2_ - C1_

	// (10)
	var dh float64

	if C1_*C2_ == 0.0 {
		dh = 0.0
	} else {
		switch {
		case math.Abs(h2_-h1_) <= 180.0:
			dh = h2_ - h1_
		case h2_-h1_ > 180.0:
			dh = h2_ - h1_ - 360.0
		case h2_-h1_ < -180.0:
			dh = h2_ - h1_ + 360.0
		}
	}

	// (11)
	dH := 2.0 * math.Sqrt(C1_*C2_) * math.Sin(rad(dh/2.0)) // (11)

	// 3 - Calculate CIEDE2000
	//

	// (12)
	Lb := (l1 + l2) / 2.0

	// (13)
	Cb_ := (C1_ + C2_) / 2.0
	Cb7 := math.Pow(Cb_, 7.0)

	// (14)
	var hb_ float64

	if C1_*C2_ != 0.0 {
		switch {
		case math.Abs(h1_-h2_) <= 180.0:
			hb_ = (h1_ + h2_) / 2.0
		case math.Abs(h1_-h2_) > 180.0 && h1_+h2_ < 360.0:
			hb_ = (h1_ + h2_ + 360.0) / 2.0
		case math.Abs(h1_-h2_) > 180.0 && h1_+h2_ >= 360.0:
			hb_ = (h1_ + h2_ - 360.0) / 2.0
		}
	} else {
		hb_ = h1_ + h2_
	}

	// (15)
	T := 1.0 - 0.17*math.Cos(rad(hb_-30.0)) + 0.24*math.Cos(rad(2.0*hb_)) + 0.32*math.Cos(rad(3.0*hb_+6.0)) - 0.2*math.Cos(rad(4.0*hb_-63.0))

	// (16)
	dt := 30.0 * math.Exp(-math.Pow((hb_-275.0)/25.0, 2.0))

	// (17)
	Rc := 2.0 * math.Sqrt(Cb7/(Cb7+6103515625.0))

	// (18)
	SL := 1.0 + (0.015*math.Pow(Lb-50.0, 2.0))/math.Sqrt(20.0+math.Pow(Lb-50.0, 2.0))

	// (19)
	SC := 1.0 + 0.045*Cb_

	// (20)
	SH := 1.0 + 0.015*Cb_*T

	// (21)
	RT := -math.Sin(rad(2.0*dt)) * Rc

	// (22)
	deltaE = math.Sqrt(math.Pow(dL/(kL*SL), 2.0) + math.Pow(dC/(kC*SC), 2.0) + math.Pow(dH/(kH*SH), 2.0) + RT*(dC/(kC*SC))*(dH/(kH*SH)))

	return
}

func rad(deg float64) (r float64) {
	r = deg * math.Pi / 180.0

	return
}
