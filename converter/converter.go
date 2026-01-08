package converter

import (
	"fmt"
	"image"
	"image/color"
	"io"
)

// Converter converts images to ASCII representation.
type Converter interface {
	Convert(img image.Image, w io.Writer, width int) error
}

// ASCIIConverter converts images to ASCII art.
type ASCIIConverter struct {
	charset Charset
}

// NewASCIIConverter creates a new ASCII converter with the given charset.
func NewASCIIConverter(charset Charset) *ASCIIConverter {
	return &ASCIIConverter{charset: charset}
}

// Convert writes the ASCII representation of img to w, scaled to the given width.
func (c *ASCIIConverter) Convert(img image.Image, w io.Writer, width int) error {
	bounds := img.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()

	if imgWidth == 0 || imgHeight == 0 {
		return fmt.Errorf("image has zero dimensions")
	}

	// Calculate scaling factors
	scaleX := float64(imgWidth) / float64(width)
	scaleY := scaleX * 2 // Compensate for character aspect ratio (chars are ~2x taller than wide)

	height := int(float64(imgHeight) / scaleY)
	if height == 0 {
		height = 1
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Sample pixel at scaled position
			srcX := int(float64(x) * scaleX)
			srcY := int(float64(y) * scaleY)

			// Clamp to image bounds
			if srcX >= imgWidth {
				srcX = imgWidth - 1
			}
			if srcY >= imgHeight {
				srcY = imgHeight - 1
			}

			pixel := img.At(bounds.Min.X+srcX, bounds.Min.Y+srcY)
			gray := toGrayscale(pixel)
			char := c.charset.CharFor(gray)

			fmt.Fprint(w, string(char))
		}
		fmt.Fprintln(w)
	}

	return nil
}

// toGrayscale converts a color to its grayscale value using luminance formula.
func toGrayscale(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()
	// RGBA returns 16-bit values (0-65535), convert to 8-bit grayscale
	// Using ITU-R BT.601 luminance formula
	gray := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 256
	return uint8(gray)
}
