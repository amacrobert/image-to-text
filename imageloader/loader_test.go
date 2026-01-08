package imageloader

import (
	"os"
	"testing"
)

func TestFileLoader_Load_FileNotFound(t *testing.T) {
	loader := NewFileLoader()
	_, err := loader.Load("nonexistent.png")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}

func TestFileLoader_Load_InvalidImage(t *testing.T) {
	// Create a temporary file with invalid image data
	tmpfile, err := os.CreateTemp("", "test*.png")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	tmpfile.WriteString("not an image")
	tmpfile.Close()

	loader := NewFileLoader()
	_, err = loader.Load(tmpfile.Name())
	if err == nil {
		t.Error("expected error for invalid image data")
	}
}
