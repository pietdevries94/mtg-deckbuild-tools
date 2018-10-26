.PHONY: all
all: prep build

.PHONY: prep
prep:
	cd ../frontend; yarn build
	go generate ./...

.PHONY: build
build:
	go build -o bin/backend cmd/local/*.go