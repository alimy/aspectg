.PHONY: default clean

LDFLAGS += -X "github.com/alimy/aspectg/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/aspectg/version.GitHash=$(shell git rev-parse HEAD)"

default:
	go build -ldflags '$(LDFLAGS)' .

clean:
	rm aspectg