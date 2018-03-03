FROM golang:1.10.0-alpine3.7 as builder
LABEL maintainer="Mike Jones"

RUN mkdir -p /go/src/github.com/n7st/quoteDB
COPY . /go/src/github.com/n7st/quoteDB

RUN apk update && apk upgrade && \
    apk add --no-cache git gcc musl-dev

WORKDIR /go/src/github.com/n7st/quoteDB

RUN go get ./... && \
    go build -o /bin/quotedb cmd/quoteDB/main.go && \
    apk del --purge git gcc musl-dev && \
    rm -rf /go/bin /go/pkg /var/cache/apk/*

FROM alpine:3.7

COPY --from=builder /bin/quotedb /bin/quotedb

RUN chmod +x /bin/quotedb
RUN mkdir -p /opt/quotedb

WORKDIR /opt/quotedb

ADD ./view /opt/quotedb/view

