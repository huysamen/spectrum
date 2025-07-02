package spectrum

import "fmt"

// RGBui8ToHexString converts RGB uint8 values (0-255) to a hex string (e.g., "FFAABB")
func RGBui8ToHexString(r, g, b uint8) string {
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}

// RGBui8ToHexValue converts RGB uint8 values (0-255) to a numerical hex value (e.g., 0xFFAABB)
func RGBui8ToHexValue(r, g, b uint8) int {
	return int(r)<<16 | int(g)<<8 | int(b)
}
