package spectrum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeltaE(t *testing.T) {
	type args struct {
		lab1 [3]float64
		lab2 [3]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Black Δ Black", args: args{lab1: Black.Lab, lab2: Black.Lab}, want: 0.00000000},
		{name: "Black Δ Red", args: args{lab1: Black.Lab, lab2: Red.Lab}, want: 117.32667126},
		{name: "Black Δ Green", args: args{lab1: Black.Lab, lab2: Green.Lab}, want: 148.47163136},
		{name: "Black Δ Blue", args: args{lab1: Black.Lab, lab2: Blue.Lab}, want: 137.64652369},
		{name: "Black Δ Yellow", args: args{lab1: Black.Lab, lab2: Yellow.Lab}, want: 137.21080020},
		{name: "Black Δ Cyan", args: args{lab1: Black.Lab, lab2: Cyan.Lab}, want: 103.98986652},
		{name: "Black Δ Magenta", args: args{lab1: Black.Lab, lab2: Magenta.Lab}, want: 130.33746279},
		{name: "Black Δ White", args: args{lab1: Black.Lab, lab2: White.Lab}, want: 100.00000000},
		{name: "White Δ Black", args: args{lab1: White.Lab, lab2: Black.Lab}, want: 100.00000000},
		{name: "White Δ Red", args: args{lab1: White.Lab, lab2: Red.Lab}, want: 114.53033671},
		{name: "White Δ Green", args: args{lab1: White.Lab, lab2: Green.Lab}, want: 120.39770828},
		{name: "White Δ Blue", args: args{lab1: White.Lab, lab2: Blue.Lab}, want: 149.96274070},
		{name: "White Δ Yellow", args: args{lab1: White.Lab, lab2: Yellow.Lab}, want: 96.94285527},
		{name: "White Δ Cyan", args: args{lab1: White.Lab, lab2: Cyan.Lab}, want: 50.90311809},
		{name: "White Δ Magenta", args: args{lab1: White.Lab, lab2: Magenta.Lab}, want: 122.16464829},
		{name: "White Δ White", args: args{lab1: White.Lab, lab2: White.Lab}, want: 0.00000000},
		{name: "Red Δ Green", args: args{lab1: Red.Lab, lab2: Green.Lab}, want: 170.56559542},
		{name: "Green Δ Blue", args: args{lab1: Green.Lab, lab2: Blue.Lab}, want: 258.68020299},
		{name: "Blue Δ Yellow", args: args{lab1: Blue.Lab, lab2: Yellow.Lab}, want: 235.14452673},
		{name: "Yellow Δ Cyan", args: args{lab1: Yellow.Lab, lab2: Cyan.Lab}, want: 111.96158701},
		{name: "Cyan Δ Magenta", args: args{lab1: Cyan.Lab, lab2: Magenta.Lab}, want: 156.64933991},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DeltaE(tt.args.lab1[0], tt.args.lab1[1], tt.args.lab1[2], tt.args.lab2[0], tt.args.lab2[1], tt.args.lab2[2])

			assert.InDeltaf(
				t,
				tt.want,
				d,
				0.000001,
				"DeltaE(%f, %f, %f, %f, %f, %f) = %f, want %f",
				tt.args.lab1[0], tt.args.lab1[1], tt.args.lab1[2], tt.args.lab2[0], tt.args.lab2[1], tt.args.lab2[2], d, tt.want,
			)
		})
	}
}

func TestDeltaE94(t *testing.T) {
	type args struct {
		lab1 [3]float64
		lab2 [3]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Black Δ Black", args: args{lab1: Black.Lab, lab2: Black.Lab}, want: 0.00000000},
		{name: "Black Δ Red", args: args{lab1: Black.Lab, lab2: Red.Lab}, want: 117.32667126},
		{name: "Black Δ Green", args: args{lab1: Black.Lab, lab2: Green.Lab}, want: 148.47163136},
		{name: "Black Δ Blue", args: args{lab1: Black.Lab, lab2: Blue.Lab}, want: 137.64652369},
		{name: "Black Δ Yellow", args: args{lab1: Black.Lab, lab2: Yellow.Lab}, want: 137.21080020},
		{name: "Black Δ Cyan", args: args{lab1: Black.Lab, lab2: Cyan.Lab}, want: 103.98986652},
		{name: "Black Δ Magenta", args: args{lab1: Black.Lab, lab2: Magenta.Lab}, want: 130.33746279},
		{name: "Black Δ White", args: args{lab1: Black.Lab, lab2: White.Lab}, want: 100.00000000},
		{name: "White Δ Black", args: args{lab1: White.Lab, lab2: Black.Lab}, want: 100.00000000},
		{name: "White Δ Red", args: args{lab1: White.Lab, lab2: Red.Lab}, want: 114.50774760},
		{name: "White Δ Green", args: args{lab1: White.Lab, lab2: Green.Lab}, want: 120.36950583},
		{name: "White Δ Blue", args: args{lab1: White.Lab, lab2: Blue.Lab}, want: 149.93448313},
		{name: "White Δ Yellow", args: args{lab1: White.Lab, lab2: Yellow.Lab}, want: 96.91992879},
		{name: "White Δ Cyan", args: args{lab1: White.Lab, lab2: Cyan.Lab}, want: 50.89143783},
		{name: "White Δ Magenta", args: args{lab1: White.Lab, lab2: Magenta.Lab}, want: 122.13878514},
		{name: "White Δ White", args: args{lab1: White.Lab, lab2: White.Lab}, want: 0.00000000},
		{name: "Red Δ Green", args: args{lab1: Red.Lab, lab2: Green.Lab}, want: 73.43089410},
		{name: "Green Δ Blue", args: args{lab1: Green.Lab, lab2: Blue.Lab}, want: 105.90485427},
		{name: "Blue Δ Yellow", args: args{lab1: Blue.Lab, lab2: Yellow.Lab}, want: 98.64892532},
		{name: "Yellow Δ Cyan", args: args{lab1: Yellow.Lab, lab2: Cyan.Lab}, want: 42.72170021},
		{name: "Cyan Δ Magenta", args: args{lab1: Cyan.Lab, lab2: Magenta.Lab}, want: 87.43264482},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DeltaE94(tt.args.lab1[0], tt.args.lab1[1], tt.args.lab1[2], tt.args.lab2[0], tt.args.lab2[1], tt.args.lab2[2])

			assert.InDeltaf(
				t,
				tt.want,
				d,
				0.000001,
				"DeltaE94(%f, %f, %f, %f, %f, %f) = %f, want %f",
				tt.args.lab1[0], tt.args.lab1[1], tt.args.lab1[2], tt.args.lab2[0], tt.args.lab2[1], tt.args.lab2[2], d, tt.want,
			)
		})
	}
}

func TestDeltaE2000(t *testing.T) {
	type args struct {
		lab1 [3]float64
		lab2 [3]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// CIEDE2000 Test Data
		// http://www.ece.rochester.edu/~gsharma/ciede2000/ciede2000noteCRNA.pdf
		{name: "CIEDE2000 Pair 1", args: args{lab1: [3]float64{50.0000, 2.6772, -79.7751}, lab2: [3]float64{50.0000, 0.0000, -82.7485}}, want: 2.0425},
		{name: "CIEDE2000 Pair 2", args: args{lab1: [3]float64{50.0000, 3.1571, -77.2803}, lab2: [3]float64{50.0000, 0.0000, -82.7485}}, want: 2.8615},
		{name: "CIEDE2000 Pair 3", args: args{lab1: [3]float64{50.0000, 2.8361, -74.0200}, lab2: [3]float64{50.0000, 0.0000, -82.7485}}, want: 3.4412},
		{name: "CIEDE2000 Pair 4", args: args{lab1: [3]float64{50.0000, -1.3802, -84.2814}, lab2: [3]float64{50.0000, 0.0000, -82.7485}}, want: 1.0000},
		{name: "CIEDE2000 Pair 5", args: args{lab1: [3]float64{50.0000, -1.1848, -84.8006}, lab2: [3]float64{50.0000, 0.0000, -82.7485}}, want: 1.0000},
		{name: "CIEDE2000 Pair 6", args: args{lab1: [3]float64{50.0000, -0.9009, -85.5211}, lab2: [3]float64{50.0000, 0.0000, -82.7485}}, want: 1.0000},
		{name: "CIEDE2000 Pair 7", args: args{lab1: [3]float64{50.0000, 0.0000, 0.0000}, lab2: [3]float64{50.0000, -1.0000, 2.0000}}, want: 2.3669},
		{name: "CIEDE2000 Pair 8", args: args{lab1: [3]float64{50.0000, -1.0000, 2.0000}, lab2: [3]float64{50.0000, 0.0000, 0.0000}}, want: 2.3669},
		{name: "CIEDE2000 Pair 9", args: args{lab1: [3]float64{50.0000, 2.4900, -0.0010}, lab2: [3]float64{50.0000, -2.4900, 0.0009}}, want: 7.1792},
		{name: "CIEDE2000 Pair 10", args: args{lab1: [3]float64{50.0000, 2.4900, -0.0010}, lab2: [3]float64{50.0000, -2.4900, 0.0010}}, want: 7.1792},
		{name: "CIEDE2000 Pair 11", args: args{lab1: [3]float64{50.0000, 2.4900, -0.0010}, lab2: [3]float64{50.0000, -2.4900, 0.0011}}, want: 7.2195},
		{name: "CIEDE2000 Pair 12", args: args{lab1: [3]float64{50.0000, 2.4900, -0.0010}, lab2: [3]float64{50.0000, -2.4900, 0.0012}}, want: 7.2195},
		{name: "CIEDE2000 Pair 13", args: args{lab1: [3]float64{50.0000, -0.0010, 2.4900}, lab2: [3]float64{50.0000, 0.0009, -2.4900}}, want: 4.8045},
		{name: "CIEDE2000 Pair 14", args: args{lab1: [3]float64{50.0000, -0.0010, 2.4900}, lab2: [3]float64{50.0000, 0.0010, -2.4900}}, want: 4.8045},
		{name: "CIEDE2000 Pair 15", args: args{lab1: [3]float64{50.0000, -0.0010, 2.4900}, lab2: [3]float64{50.0000, 0.0011, -2.4900}}, want: 4.7461},
		{name: "CIEDE2000 Pair 16", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{50.0000, 0.0000, -2.5000}}, want: 4.3065},
		{name: "CIEDE2000 Pair 17", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{73.0000, 25.0000, -18.0000}}, want: 27.1492},
		{name: "CIEDE2000 Pair 18", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{61.0000, -5.0000, 29.0000}}, want: 22.8977},
		{name: "CIEDE2000 Pair 19", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{56.0000, -27.0000, -3.0000}}, want: 31.9030},
		{name: "CIEDE2000 Pair 20", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{58.0000, 24.0000, 15.0000}}, want: 19.4535},
		{name: "CIEDE2000 Pair 21", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{50.0000, 3.1736, 0.5854}}, want: 1.0000},
		{name: "CIEDE2000 Pair 22", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{50.0000, 3.2972, 0.0000}}, want: 1.0000},
		{name: "CIEDE2000 Pair 23", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{50.0000, 1.8634, 0.5757}}, want: 1.0000},
		{name: "CIEDE2000 Pair 24", args: args{lab1: [3]float64{50.0000, 2.5000, 0.0000}, lab2: [3]float64{50.0000, 3.2592, 0.3350}}, want: 1.0000},
		{name: "CIEDE2000 Pair 25", args: args{lab1: [3]float64{60.2574, -34.0099, 36.2677}, lab2: [3]float64{60.4626, -34.1751, 39.4387}}, want: 1.2644},
		{name: "CIEDE2000 Pair 26", args: args{lab1: [3]float64{63.0109, -31.0961, -5.8663}, lab2: [3]float64{62.8187, -29.7946, -4.0864}}, want: 1.2630},
		{name: "CIEDE2000 Pair 27", args: args{lab1: [3]float64{61.2901, 3.7196, -5.3901}, lab2: [3]float64{61.4292, 2.2480, -4.9620}}, want: 1.8731},
		{name: "CIEDE2000 Pair 28", args: args{lab1: [3]float64{35.0831, -44.1164, 3.7933}, lab2: [3]float64{35.0232, -40.0716, 1.5901}}, want: 1.8645},
		{name: "CIEDE2000 Pair 29", args: args{lab1: [3]float64{22.7233, 20.0904, -46.6940}, lab2: [3]float64{23.0331, 14.9730, -42.5619}}, want: 2.0373},
		{name: "CIEDE2000 Pair 30", args: args{lab1: [3]float64{36.4612, 47.8580, 18.3852}, lab2: [3]float64{36.2715, 50.5065, 21.2231}}, want: 1.4146},
		{name: "CIEDE2000 Pair 31", args: args{lab1: [3]float64{90.8027, -2.0831, 1.4410}, lab2: [3]float64{91.1528, -1.6435, 0.0447}}, want: 1.4441},
		{name: "CIEDE2000 Pair 32", args: args{lab1: [3]float64{90.9257, -0.5406, -0.9208}, lab2: [3]float64{88.6381, -0.8985, -0.7239}}, want: 1.5381},
		{name: "CIEDE2000 Pair 33", args: args{lab1: [3]float64{6.7747, -0.2908, -2.4247}, lab2: [3]float64{5.8714, -0.0985, -2.2286}}, want: 0.6377},
		{name: "CIEDE2000 Pair 34", args: args{lab1: [3]float64{2.0776, 0.0795, -1.1350}, lab2: [3]float64{0.9033, -0.0636, -0.5514}}, want: 0.9082},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1 := DeltaE2000(tt.args.lab1[0], tt.args.lab1[1], tt.args.lab1[2], tt.args.lab2[0], tt.args.lab2[1], tt.args.lab2[2])
			d2 := DeltaE2000(tt.args.lab2[0], tt.args.lab2[1], tt.args.lab2[2], tt.args.lab1[0], tt.args.lab1[1], tt.args.lab1[2])

			assert.InDeltaf(t, d1, d2, 0.05, "DeltaE2000(%f, %f, %f, %f, %f, %f) = %f, want %f", tt.args.lab1[0], tt.args.lab1[1], tt.args.lab1[2], tt.args.lab2[0], tt.args.lab2[1], tt.args.lab2[2], d1, d2)
			assert.InDeltaf(
				t,
				tt.want,
				d1,
				0.0001,
				"DeltaE2000(%f, %f, %f, %f, %f, %f) = %f, want %f",
				tt.args.lab1[0], tt.args.lab1[1], tt.args.lab1[2], tt.args.lab2[0], tt.args.lab2[1], tt.args.lab2[2], d1, tt.want,
			)
		})
	}
}
