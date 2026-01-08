.PHONY: build clean test

build:
	@mkdir -p dist
	go build -o dist/image-to-text .

test:
	go test ./...

clean:
	rm -rf dist
