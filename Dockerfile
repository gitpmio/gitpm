FROM golang:latest

RUN go get github.com/astaxie/beego && go get github.com/beego/bee

WORKDIR /go/src/github.com/gitpmio/gitpm/gitpm-app
CMD ["bee", "run"]
