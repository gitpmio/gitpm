FROM golang:latest

RUN mkdir -p /go/pkg/
RUN mkdir -p /go/bin/
RUN mkdir -p /go/src/github.com/gitpmio/gitpm

RUN go get github.com/revel/cmd/revel
RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/gitpmio/gitpm/gitpm-app
RUN dep ensure

ENTRYPOINT "revel run github.com/gitpmio/gitpm/gitpm-app"
