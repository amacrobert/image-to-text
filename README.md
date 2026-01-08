# image-to-text

A CLI tool that converts images to ASCII art.

## Build

```bash
go build .
```

## Usage

```bash
./image-to-text [options] <image-file>
```

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `-width` | 80 | Output width in characters |
| `-invert` | false | Invert brightness (for dark terminal backgrounds) |

### Examples

```bash
# Basic usage
./image-to-text photo.png

# Wider output
./image-to-text -width 120 photo.jpg

# For dark terminal backgrounds
./image-to-text -invert photo.png

# Run without building
go run . photo.png
```

## Supported Formats

- PNG
- JPEG
- GIF
