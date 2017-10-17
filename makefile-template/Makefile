# This is based on Radomir Sohlich
# https://sohlich.github.io/post/go_makefile/

# Go parameters
GOBUILD				= go build
GOCLEAN				= go clean
GOTEST				= go test
BINARY_NAME		= $(shell pwd | sed 's:.*/::')
BINARY_UNIX		= $(BINARY_NAME)_unix
PID           = .pid
GLIDE					:= $(shell command -v glide 2> /dev/null)

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -rf vendor

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

kill:
	@kill `cat $(PID)` || true

install:
	@echo "Checking for glide"

ifdef GLIDE
	@glide install
else
	$(warning "Skipping installation of Go dependencies: glide is not installed")
endif
