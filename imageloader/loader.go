package imageloader

import (
	"fmt"
	"image"
	_ "image/gif"  // Register GIF decoder
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"os"
)

// Loader loads images from various sources.
type Loader interface {
	Load(path string) (image.Image, error)
}

// FileLoader loads images from the filesystem.
type FileLoader struct{}

// NewFileLoader creates a new FileLoader.
func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

// Load reads and decodes an image from the given file path.
func (l *FileLoader) Load(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", path, err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("decoding image: %w", err)
	}

	return img, nil
}
