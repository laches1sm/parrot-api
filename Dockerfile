FROM golang:1.11-alpine3.8

RUN apk update && apk upgrade && apk add --no-cache \
        bash \
        git  \
        libcurl \
        curl \
        build-base \
        py-pip

RUN pip install docker-compose 

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin v1.11.2

RUN go get -u goa.design/goa/... && \
    mkdir /build

EXPOSE 7000

