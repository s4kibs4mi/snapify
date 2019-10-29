FROM golang:stretch AS builder

RUN apt-get update && apt-get install ca-certificates && rm -rf /var/cache/apt/*
#RUN apt-get install git openssh

ENV GOPATH=/go

ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GO111MODULE=on

COPY . $GOPATH/src/github.com/s4kibs4mi/snapify
WORKDIR $GOPATH/src/github.com/s4kibs4mi/snapify

#RUN go get github.com/ugorji/go@v1.1.2-0.20180831062425-e253f1f20942

RUN go get .
#RUN rm /go/pkg/mod/github.com/coreos/etcd@v3.3.10+incompatible/client/keys.generated.go
#RUN cp ./hacks/keys.generated.go /go/pkg/mod/github.com/coreos/etcd@v3.3.10+incompatible/client/

RUN go build -v -o snapify
RUN mv snapify /go/bin/snapify

FROM chromedp/headless-shell:latest
WORKDIR /root

COPY --from=builder /go/bin/snapify /usr/local/bin/snapify

ENTRYPOINT ["snapify"]
