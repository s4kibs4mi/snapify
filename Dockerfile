FROM --platform=linux/amd64 golang:1.18-alpine AS builder

ENV GOPATH=/go

ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GO111MODULE=on

COPY . $GOPATH/src/github.com/s4kibs4mi/snapify
WORKDIR $GOPATH/src/github.com/s4kibs4mi/snapify

RUN go mod tidy

RUN go build -v -o snapify-app ./cmd/app
RUN go build -v -o snapify-worker ./cmd/worker
RUN mv snapify-app /go/bin/snapify-app
RUN mv snapify-worker /go/bin/snapify-worker

FROM --platform=linux/amd64 alpine

WORKDIR /root

COPY --from=builder /go/bin/snapify-app /usr/local/bin/snapify-app
COPY --from=builder /go/bin/snapify-worker /usr/local/bin/snapify-worker
