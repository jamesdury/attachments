BINARY_NAME=attachments

all: test build

.PHONY: build
build:
	npm run icons
	npm run build
	mkdir -p build
	go build -o ./build/${BINARY_NAME} main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -r node_modules
	rm build/${BINARY_NAME}
	go clean
