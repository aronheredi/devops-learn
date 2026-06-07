# Variables
BINARY_NAME=statuspage-app
DOCKER_IMAGE=statuspage-app
VERSION=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

# Go linker flags to inject SHA and Timestamp
LDFLAGS=-ldflags "-X main.CommitSHA=$(VERSION) -X main.buildDate=$(BUILD_TIME)"

.PHONY: all build test image run clean

all: test build

build:
	@echo "Building local binary..."
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(BINARY_NAME) ./...

test:
	@echo "Running tests..."
	go test ./... -v

image:
	@echo "Building secure Docker image [$(DOCKER_IMAGE):$(VERSION)]..."
	docker build \
		--build-arg "GIT_SHA=$(VERSION)" \
		--build-arg "BUILD_TIME=$(BUILD_TIME)" \
		-t $(DOCKER_IMAGE):$(VERSION) \
		-t $(DOCKER_IMAGE):latest .

run: image
	@echo "Running container locally..."
	docker run --rm -p 8080:8080 $(DOCKER_IMAGE):$(VERSION)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)