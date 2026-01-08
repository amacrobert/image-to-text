# image-to-text

A CLI tool that converts images to ASCII art.

## Build

```bash
make build
```

This creates the binary at `dist/image-to-text`.

Other targets:
- `make test` - run tests
- `make clean` - remove build artifacts

## Usage

```bash
./dist/image-to-text [options] <image-file>
```

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `-width` | 80 | Output width in characters |
| `-invert` | false | Invert brightness (for dark terminal backgrounds) |

### Examples

```bash
# Basic usage
./dist/image-to-text photo.png

# Wider output
./dist/image-to-text -width 120 photo.jpg

# For dark terminal backgrounds
./dist/image-to-text -invert photo.png

# Run without building
go run . photo.png
```

## Supported Formats

- PNG
- JPEG
- GIF
