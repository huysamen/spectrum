package spectrum

type Colour struct {
	Name    string
	RGB     [3]float64
	RGBui8  [3]uint8
	RGBui32 uint32
	XYZ     [3]float64
	Lab     [3]float64
	HSL     [3]float64
}

var (
	Black = Colour{
		Name:    "Black",
		RGB:     [3]float64{0.00000000, 0.00000000, 0.00000000},
		RGBui8:  [3]uint8{0, 0, 0},
		RGBui32: 0x000000,
		XYZ:     [3]float64{0.00000000, 0.00000000, 0.00000000},
		Lab:     [3]float64{0.00000000, 0.00000000, 0.00000000},
		HSL:     [3]float64{0.00000000, 0.00000000, 0.00000000},
	}

	Red = Colour{
		Name:    "Red",
		RGB:     [3]float64{1.00000000, 0.00000000, 0.00000000},
		RGBui8:  [3]uint8{255, 0, 0},
		RGBui32: 0xff0000,
		XYZ:     [3]float64{0.41245300, 0.21267100, 0.01933400},
		Lab:     [3]float64{53.240587944, 80.092308226, 67.20275104},
		HSL:     [3]float64{0.00000000, 1.00000000, 0.50000000},
	}

	Green = Colour{
		Name:    "Green",
		RGB:     [3]float64{0.00000000, 1.00000000, 0.00000000},
		RGBui8:  [3]uint8{0, 255, 0},
		RGBui32: 0x00ff00,
		XYZ:     [3]float64{0.35758000, 0.71516000, 0.11919300},
		Lab:     [3]float64{87.73509949, -86.18302974, 83.17970318},
		HSL:     [3]float64{0.33333333, 1.00000000, 0.50000000},
	}

	Blue = Colour{
		Name:    "Blue",
		RGB:     [3]float64{0.00000000, 0.00000000, 1.00000000},
		RGBui8:  [3]uint8{0, 0, 255},
		RGBui32: 0x0000ff,
		XYZ:     [3]float64{0.18042300, 0.07216900, 0.95022700},
		Lab:     [3]float64{32.29567257, 79.18559091, -107.85730021},
		HSL:     [3]float64{0.66666667, 1.00000000, 0.50000000},
	}

	Yellow = Colour{
		Name:    "Yellow",
		RGB:     [3]float64{1.00000000, 1.00000000, 0.00000000},
		RGBui8:  [3]uint8{255, 255, 0},
		RGBui32: 0xffff00,
		XYZ:     [3]float64{0.77003300, 0.92783100, 0.13852700},
		Lab:     [3]float64{97.13950704, -21.55468102, 94.47812228},
		HSL:     [3]float64{0.16666667, 1.00000000, 0.50000000},
	}

	Cyan = Colour{
		Name:    "Cyan",
		RGB:     [3]float64{0.00000000, 1.00000000, 1.00000000},
		RGBui8:  [3]uint8{0, 255, 255},
		RGBui32: 0x00ffff,
		XYZ:     [3]float64{0.53800300, 0.78732900, 1.06942000},
		Lab:     [3]float64{91.11330144, -48.09059623, -14.12632982},
		HSL:     [3]float64{0.50000000, 1.00000000, 0.50000000},
	}

	Magenta = Colour{
		Name:    "Magenta",
		RGB:     [3]float64{1.00000000, 0.00000000, 1.00000000},
		RGBui8:  [3]uint8{255, 0, 255},
		RGBui32: 0xff00ff,
		XYZ:     [3]float64{0.59287600, 0.28484000, 0.96956100},
		Lab:     [3]float64{60.32350653, 98.23305386, -60.82101524},
		HSL:     [3]float64{0.83333333, 1.00000000, 0.50000000},
	}

	White = Colour{
		Name:    "White",
		RGB:     [3]float64{1.00000000, 1.00000000, 1.00000000},
		RGBui8:  [3]uint8{255, 255, 255},
		RGBui32: 0xffffff,
		XYZ:     [3]float64{0.95045600, 1.00000000, 1.08875400},
		Lab:     [3]float64{100.00000000, -0.00245494, 0.00465342},
		HSL:     [3]float64{0.00000000, 0.00000000, 1.00000000},
	}

	AllColours = []Colour{
		Black,
		Red,
		Green,
		Blue,
		Yellow,
		Cyan,
		Magenta,
		White,
	}
)