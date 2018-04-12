GO := GOPATH=$(shell go env GOROOT)/bin:"$(shell pwd)" GOOS=$(OS) GOARCH=$(ARCH) go
#GO := GOPATH=$(shell go env GOROOT)/bin:"$(shell pwd)" go
GOGETTER := GOPATH="$(shell pwd)" GOOS=$(OS) GOARCH=$(ARCH) go get -u
.PHONY: all bloxtool-go

all: clean go_get_deps bloxtool-go

bloxtool-go:
	$(GO) build $(GOOPTS) -o bin/bloxtool-go main.go get_config.go host_action.go

go_get_deps:
	$(GOGETTER) gopkg.in/ini.v1
	$(GOGETTER) github.com/docopt/docopt-go 
	$(GOGETTER) github.com/rtucker-mozilla/go-infoblox

test:
	$(GO) test ./tests/

tests:
	$(GO) test ./tests/

clean:
	rm -rf bin src/github.com src/bitbucket.org src/code.google.com src/golang.org src/gopkg.in
