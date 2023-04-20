BINARY_NAME=attachments

default: test build

build:
	npm run icons
	npm run sass
	npm run css
	mkdir -p build
	go build -o ./build/${BINARY_NAME} main.go

test:
	go test -v ./...

clean:
	rm -r node_modules
	rm build/${BINARY_NAME}
	go clean

.PHONY: build test clean
