package spectrum

import "fmt"

// HexStringToRGBui8 converts a hex string (e.g., "FFAABB" or "#FFAABB") to RGB uint8 values.
func HexStringToRGBui8(hex string) (r, g, b uint8, err error) {
	if len(hex) == 7 && hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex string length: %s", hex)
	}

	var rv, gv, bv int

	_, err = fmt.Sscanf(hex, "%02X%02X%02X", &rv, &gv, &bv)
	if err != nil {
		_, err = fmt.Sscanf(hex, "%02x%02x%02x", &rv, &gv, &bv)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid hex string: %s", hex)
		}
	}

	return uint8(rv), uint8(gv), uint8(bv), nil
}

// HexValueToRGBui8 converts a numerical hex value (e.g., 0xFFAABB) to RGB uint8 values.
func HexValueToRGBui8(hex int) (r, g, b uint8) {
	r = uint8((hex >> 16) & 0xFF)
	g = uint8((hex >> 8) & 0xFF)
	b = uint8(hex & 0xFF)

	return
}
