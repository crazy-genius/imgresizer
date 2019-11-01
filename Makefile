pkgs=$(shell go list ./... | fgrep -v /vendor)
files=$(shell find . -type f -name "*.go" -print | fgrep -v /vendor)

default: build

clean:
	rm -rf bin
create_folders:
	mkdir bin
fmt:
	find . -type f -name "*.go" -print | xargs gofmt -w
lint:
	go get golang.org/x/lint/golint
	go get honnef.co/go/tools/cmd/staticcheck
	go get github.com/kisielk/errcheck
	go get golang.org/x/tools/cmd/goimports

	golint -set_exit_status $(pkgs)
	go vet $(pkgs)
	staticcheck $(pkgs)
	errcheck $(pkgs)
	gocritic check $(pkgs)
	goimports -l $(files)
 

build: fmt lint clean create_folders 
	go build -o bin/imgresizer ./cmd/imgresizer/main.go

install_critic:
	go get github.com/go-lintpack/lintpack/...
	go get github.com/go-critic/go-critic/...
	lintpack build -o $(GOPATH)/bin/gocritic -linter.version='v0.3.4' -linter.name='gocritic' github.com/go-critic/go-critic/checkers
