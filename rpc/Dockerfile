FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
COPY . .
RUN go build -ldflags="-s -w" -o /app/bitMapRpcServ /build/zero/bitmaprpc.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Jakarta

WORKDIR /app
COPY --from=builder /app/bitMapRpcServ /app/bitMapRpcServ

EXPOSE 80 8081
VOLUME /app/tmp/db

CMD ["./bitMapRpcServ"]
