# FROM golang:latest

# ENV GOPROXY https://goproxy.cn,direct
# WORKDIR $GOPATH/src/b
# COPY . $GOPATH/src/b
# RUN go build .

# EXPOSE 8000
# ENTRYPOINT ["./b"]

FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
COPY . .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o b main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/b /build/b
COPY --from=0 /build/conf/app.ini ./conf/

EXPOSE 8000
CMD ["./b"]
