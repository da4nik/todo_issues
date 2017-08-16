BINARY=todo_issues
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-linkmode external -s -w -extldflags -static -X github.com/da4nik/todo_issues/main.Version=${VERSION} -X github.com/da4nik/todo_issues/main.BuildTime=${BUILD_TIME}"

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

.PHONY: build run install clean
.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	# glide install --skip-test
	go build ${LDFLAGS} -o ${BINARY} todo_issues.go

build: $(BINARY)

run:
	@go run ${BINARY}.go

install:
	go install ${LDFLAGS} ./...

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
