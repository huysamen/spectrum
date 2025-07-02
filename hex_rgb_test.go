package spectrum

import "testing"

func TestHexStringToRGBui8(t *testing.T) {
	tests := []struct {
		hex      string
		r, g, b  uint8
		hasError bool
	}{
		{"FFAABB", 255, 170, 187, false},
		{"#FFAABB", 255, 170, 187, false},
		{"000000", 0, 0, 0, false},
		{"#000000", 0, 0, 0, false},
		{"FFFFFF", 255, 255, 255, false},
		{"#FFFFFF", 255, 255, 255, false},
		{"102030", 16, 32, 48, false},
		{"#102030", 16, 32, 48, false},
		{"#FFF", 0, 0, 0, true},   // invalid length
		{"GGHHII", 0, 0, 0, true}, // invalid hex
	}
	for _, tt := range tests {
		r, g, b, err := HexStringToRGBui8(tt.hex)
		if tt.hasError {
			if err == nil {
				t.Errorf("HexStringToRGBui8(%q) expected error, got none", tt.hex)
			}
			continue
		}

		if err != nil {
			t.Errorf("HexStringToRGBui8(%q) unexpected error: %v", tt.hex, err)
			continue
		}

		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf("HexStringToRGBui8(%q) = (%d, %d, %d), want (%d, %d, %d)", tt.hex, r, g, b, tt.r, tt.g, tt.b)
		}
	}
}

func TestHexValueToRGBui8(t *testing.T) {
	tests := []struct {
		hex     int
		r, g, b uint8
	}{
		{0xFFAABB, 255, 170, 187},
		{0x000000, 0, 0, 0},
		{0xFFFFFF, 255, 255, 255},
		{0x102030, 16, 32, 48},
	}
	for _, tt := range tests {
		r, g, b := HexValueToRGBui8(tt.hex)

		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf("HexValueToRGBui8(%#06x) = (%d, %d, %d), want (%d, %d, %d)", tt.hex, r, g, b, tt.r, tt.g, tt.b)
		}
	}
}
