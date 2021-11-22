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
RUN go build -ldflags="-s -w" -o /app/bitMapServ /build/zero/api/bitmapserv.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Jakarta

WORKDIR /app
COPY --from=builder /app/bitMapServ /app/bitMapServ

EXPOSE 8080
VOLUME /app/tmp/db

CMD ["./bitMapServ"]
