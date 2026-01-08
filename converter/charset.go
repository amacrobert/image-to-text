package converter

// Charset maps grayscale values to ASCII characters.
type Charset interface {
	CharFor(grayValue uint8) rune
}

type simpleCharset struct {
	chars      string
	invert     bool
	blackpoint uint8
	whitepoint uint8
}

// NewSimpleCharset creates a charset with basic characters ordered by visual density.
// If invert is true, the mapping is reversed (for dark terminal backgrounds).
// Blackpoint and whitepoint adjust the input range: values at or below blackpoint
// become pure black, values at or above whitepoint become pure white.
func NewSimpleCharset(invert bool, blackpoint, whitepoint uint8) Charset {
	return &simpleCharset{
		chars:      "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ",
		invert:     invert,
		blackpoint: blackpoint,
		whitepoint: whitepoint,
	}
}

func (c *simpleCharset) CharFor(grayValue uint8) rune {
	// Apply blackpoint/whitepoint remapping
	if c.blackpoint < c.whitepoint {
		if grayValue <= c.blackpoint {
			grayValue = 0
		} else if grayValue >= c.whitepoint {
			grayValue = 255
		} else {
			// Scale linearly between blackpoint and whitepoint
			grayValue = uint8((int(grayValue) - int(c.blackpoint)) * 255 / (int(c.whitepoint) - int(c.blackpoint)))
		}
	}

	if c.invert {
		grayValue = 255 - grayValue
	}
	index := int(grayValue) * (len(c.chars) - 1) / 255
	return rune(c.chars[index])
}
