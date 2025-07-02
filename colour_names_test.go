package spectrum

import (
	"fmt"
	"strings"
	"testing"
)

func hexToRGBu8(hex string) (uint8, uint8, uint8) {
	var r, g, b uint8

	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	if err != nil {
		return 0, 0, 0
	}

	return r, g, b
}

func TestClosestColourNameHex(t *testing.T) {
	tests := []struct {
		name     string
		hex      string
		expected string
	}{
		{"White", "#ffffff", "White"},
		{"Black", "#000000", "Black"},
		{"Red", "#e5 00 00", "Red"},
		{"Green", "#15 b0 1a", "Green"},
		{"Blue", "#03 43 df", "Blue"},
		{"Magenta", "#c2 00 78", "Magenta"},
		{"Yellow", "#ff ff 14", "Yellow"},
		{"Brown", "#65 37 00", "Brown"},
		{"Pink", "#ff 81 c0", "Pink"},
		{"Light Blue", "#95 d0 fc", "Light Blue"},
		{"Pale Lime", "#C2F970", "Pale Lime"},
		{"Charcoal Grey", "#44344F", "Charcoal Grey"},
		{"Dusk", "#564D80", "Dusk"},
		{"Light Grey Blue", "#98A6D4", "Light Grey Blue"},
		{"Very Light Green", "#D3FCD5", "Very Light Green"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := hexToRGBu8(strings.ReplaceAll(tt.hex, " ", ""))
			result := ClosestColourName(r, g, b)

			if result != tt.expected {
				t.Errorf("ClosestColourName(hex %q) = %q, want %q", tt.hex, result, tt.expected)
			}
		})
	}
}

func TestClosestColourName(t *testing.T) {
	tests := []struct {
		name     string
		r, g, b  uint8
		expected string
	}{
		{"White", 255, 255, 255, "White"},
		{"Black", 0, 0, 0, "Black"},
		{"Red", 229, 0, 0, "Red"},
		{"Green", 21, 176, 26, "Green"},
		{"Blue", 3, 67, 223, "Blue"},
		{"Magenta", 194, 0, 120, "Magenta"},
		{"Yellow", 255, 255, 20, "Yellow"},
		{"Brown", 101, 55, 0, "Brown"},
		{"Pink", 255, 129, 192, "Pink"},
		{"Light Blue", 149, 208, 252, "Light Blue"},
		{"Chocolate", 210, 106, 30, "Chocolate"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClosestColourName(tt.r, tt.g, tt.b)
			if result != tt.expected {
				t.Errorf("ClosestColourName(%d, %d, %d) = %q, want %q", tt.r, tt.g, tt.b, result, tt.expected)
			}
		})
	}
}
