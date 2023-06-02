FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/b
COPY . $GOPATH/src/b
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./b"]