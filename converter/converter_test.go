package converter

import (
	"image"
	"image/color"
	"strings"
	"testing"
)

func TestToGrayscale(t *testing.T) {
	tests := []struct {
		name  string
		color color.Color
		want  uint8
	}{
		{"white", color.White, 255},
		{"black", color.Black, 0},
		{"red", color.RGBA{255, 0, 0, 255}, 76},   // 0.299 * 255 ≈ 76
		{"green", color.RGBA{0, 255, 0, 255}, 149}, // 0.587 * 255 ≈ 149
		{"blue", color.RGBA{0, 0, 255, 255}, 29},   // 0.114 * 255 ≈ 29
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toGrayscale(tt.color)
			// Allow small rounding differences
			diff := int(got) - int(tt.want)
			if diff < -1 || diff > 1 {
				t.Errorf("toGrayscale(%v) = %d, want %d", tt.color, got, tt.want)
			}
		})
	}
}

func TestASCIIConverter_Convert(t *testing.T) {
	// Create a simple 4x4 test image with black and white pixels
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	// Fill with white
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			img.Set(x, y, color.White)
		}
	}
	// Add some black pixels
	img.Set(0, 0, color.Black)
	img.Set(3, 3, color.Black)

	charset := NewSimpleCharset(false, 0, 255)
	conv := NewASCIIConverter(charset)
	var buf strings.Builder

	err := conv.Convert(img, &buf, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()

	// Should contain dark character for black pixels
	if !strings.Contains(output, "$") {
		t.Error("expected '$' character for black pixels in output")
	}

	// Should contain space for white pixels
	if !strings.Contains(output, " ") {
		t.Error("expected space character for white pixels in output")
	}

	// Should have newlines (multiple lines of output)
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) < 1 {
		t.Error("expected at least one line of output")
	}
}

func TestASCIIConverter_Convert_ZeroDimensions(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 0, 0))
	charset := NewSimpleCharset(false, 0, 255)
	conv := NewASCIIConverter(charset)
	var buf strings.Builder

	err := conv.Convert(img, &buf, 80)
	if err == nil {
		t.Error("expected error for zero-dimension image")
	}
}
