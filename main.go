package main

import (
	"flag"
	"fmt"
	"os"

	"image-to-text/converter"
	"image-to-text/imageloader"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	width := flag.Int("width", 80, "output width in characters")
	invert := flag.Bool("invert", false, "invert brightness (for dark terminal backgrounds)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <image-file>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Convert an image to ASCII art.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		return fmt.Errorf("expected exactly one image file argument")
	}

	imagePath := args[0]

	// Load image
	loader := imageloader.NewFileLoader()
	img, err := loader.Load(imagePath)
	if err != nil {
		return err
	}

	// Convert to ASCII
	charset := converter.NewSimpleCharset(*invert)
	conv := converter.NewASCIIConverter(charset)

	return conv.Convert(img, os.Stdout, *width)
}
