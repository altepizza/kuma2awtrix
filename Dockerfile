FROM golang:1.22 as base

LABEL org.opencontainers.image.source https://github.com/altepizza/kuma2awtrix

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src/main.go ./

RUN go build -o main .

CMD ["./main"]
