FROM golang:alpine AS builder

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add git openssh

ENV GOPATH=/go

ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GO111MODULE=on

COPY . $GOPATH/src/github.com/s4kibs4mi/snapify
WORKDIR $GOPATH/src/github.com/s4kibs4mi/snapify

RUN go get .

RUN go build -v -o snapify
RUN mv snapify /go/bin/snapify

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root

COPY --from=builder /go/bin/snapify /usr/local/bin/snapify

ENTRYPOINT ["snapify"]
