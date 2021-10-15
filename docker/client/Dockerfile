FROM tetafro/golang-gcc:1.16-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ../. ./

RUN go build -o geometry-client $PWD/cmd/client

FROM alpine:3.14
ENTRYPOINT ["geometry-client"]

COPY --from=builder /app/geometry-client /usr/local/bin/
EXPOSE 50051