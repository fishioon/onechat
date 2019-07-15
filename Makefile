VERSION=`git rev-parse --short HEAD`
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o bin/server server/*.go

proto:
	protoc -I chat/ chat/*.proto --go_out=plugins=grpc:chat \
		--js_out=import_style=commonjs:./extension/src/ \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./extension/src/

install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY:  clean install
