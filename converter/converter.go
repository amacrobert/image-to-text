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
			// Calculate the rectangle of source pixels this character represents
			srcX1 := int(float64(x) * scaleX)
			srcY1 := int(float64(y) * scaleY)
			srcX2 := int(float64(x+1) * scaleX)
			srcY2 := int(float64(y+1) * scaleY)

			// Clamp to image bounds
			if srcX2 > imgWidth {
				srcX2 = imgWidth
			}
			if srcY2 > imgHeight {
				srcY2 = imgHeight
			}

			// Ensure at least one pixel is sampled
			if srcX2 <= srcX1 {
				srcX2 = srcX1 + 1
			}
			if srcY2 <= srcY1 {
				srcY2 = srcY1 + 1
			}

			// Average all pixels in the rectangle
			var sum float64
			count := 0
			for py := srcY1; py < srcY2; py++ {
				for px := srcX1; px < srcX2; px++ {
					pixel := img.At(bounds.Min.X+px, bounds.Min.Y+py)
					sum += float64(toGrayscale(pixel))
					count++
				}
			}

			gray := uint8(sum / float64(count))
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
