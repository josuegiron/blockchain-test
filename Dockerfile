FROM golang:latest as builder

RUN mkdir /build
WORKDIR /build

COPY / .

RUN env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o main .

FROM alpine:latest

RUN mkdir -p /app

COPY --from=builder /build/main /app/

EXPOSE 9090
WORKDIR /app
CMD ["./main"]