FROM golang:1.23-alpine AS builder
LABEL authors="Guangyu He"
LABEL version="0.1"
LABEL email="me@heguangyu.net"

WORKDIR /app
COPY . .
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache tzdata
RUN ln -sf /usr/share/zoneinfo/Europe/Berlin /etc/localtime && echo "Europe/Berlin" > /etc/timezone
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8888
CMD ["./server"]
