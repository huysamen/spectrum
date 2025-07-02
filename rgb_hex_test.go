package spectrum

import "testing"

func TestRGBui8ToHexString(t *testing.T) {
	tests := []struct {
		r, g, b  uint8
		expected string
	}{
		{255, 170, 187, "FFAABB"},
		{0, 0, 0, "000000"},
		{255, 255, 255, "FFFFFF"},
		{16, 32, 48, "102030"},
	}
	for _, tt := range tests {
		result := RGBui8ToHexString(tt.r, tt.g, tt.b)
		if result != tt.expected {
			t.Errorf("RGBui8ToHexString(%d, %d, %d) = %s; want %s", tt.r, tt.g, tt.b, result, tt.expected)
		}
	}
}

func TestRGBui8ToHexValue(t *testing.T) {
	tests := []struct {
		r, g, b  uint8
		expected int
	}{
		{255, 170, 187, 0xFFAABB},
		{0, 0, 0, 0x000000},
		{255, 255, 255, 0xFFFFFF},
		{16, 32, 48, 0x102030},
	}
	for _, tt := range tests {
		result := RGBui8ToHexValue(tt.r, tt.g, tt.b)
		if result != tt.expected {
			t.Errorf("RGBui8ToHexValue(%d, %d, %d) = %#06x; want %#06x", tt.r, tt.g, tt.b, result, tt.expected)
		}
	}
}
