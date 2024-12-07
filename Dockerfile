FROM golang:1.21-alpine AS builder
LABEL authors="Guangyu He"
LABEL version="0.1"
LABEL email="me@heguangyu.net"

WORKDIR /app
COPY . .
RUN go build -o server main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8888
CMD ["./server"]
