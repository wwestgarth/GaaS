FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ../. ./

RUN go build -o geometry-server $PWD/cmd/server

FROM alpine:3.14
ENTRYPOINT ["geometry-server"]

COPY --from=builder /app/geometry-server /usr/local/bin/
EXPOSE 50051