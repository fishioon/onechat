PROJ=onechat
ORG_PATH=github.com/fishioon
REPO_PATH=$(ORG_PATH)/$(PROJ)
VERSION=`git rev-parse --short HEAD`
BUILD=`date +%FT%T%z`

$( shell mkdir -p bin ) 

LDFLAGS=-ldflags "-X $(REPO_PATH)/version.Version=${VERSION} -X $(REPO_PATH)/version.BuildTime=${BUILD}"

build:
	@go build ${LDFLAGS} -o bin/onechat $(REPO_PATH)/cmd/onechat
#	@go build ${LDFLAGS} -o bin/client $(REPO_PATH)/cmd/client

proto:
	@protoc -I/usr/local/include -I. \
		-I $(GOPATH)/src \
		--go_out=plugins=grpc:. \
		proto/*.proto

.PHONY:  clean install
