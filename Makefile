VERSION=`git rev-parse --short HEAD`
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o bin/server server/*.go
	go build ${LDFLAGS} -o bin/client client/*.go

proto:
	protoc -I chat/ chat/*.proto --go_out=plugins=grpc:chat

install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY:  clean install
