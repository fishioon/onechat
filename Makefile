VERSION=`git rev-parse --short HEAD`
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o bin/server server/*.go
	go build ${LDFLAGS} -o bin/client client/*.go

proto:
	protoc -I chat/ chat/*.proto --go_out=plugins=grpc:chat \
		--js_out=import_style=commonjs:chat \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:chat

install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY:  clean install
