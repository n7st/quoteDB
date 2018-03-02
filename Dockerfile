FROM golang:1.10.0-alpine3.7

RUN apk update && apk upgrade && \
    apk add --no-cache git gcc musl-dev

RUN go get -v github.com/n7st/quoteDB/...
