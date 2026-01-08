package converter

// Charset maps grayscale values to ASCII characters.
type Charset interface {
	CharFor(grayValue uint8) rune
}

type simpleCharset struct {
	chars  string
	invert bool
}

// NewSimpleCharset creates a charset with basic characters ordered by visual density.
// If invert is true, the mapping is reversed (for dark terminal backgrounds).
func NewSimpleCharset(invert bool) Charset {
	return &simpleCharset{
		chars:  "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ",
		invert: invert,
	}
}

func (c *simpleCharset) CharFor(grayValue uint8) rune {
	if c.invert {
		grayValue = 255 - grayValue
	}
	index := int(grayValue) * (len(c.chars) - 1) / 255
	return rune(c.chars[index])
}
