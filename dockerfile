FROM golang:1.26-alpine AS builder

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN swag init

RUN go build -o main .

FROM scratch

WORKDIR /

COPY --from=builder /app/main .
COPY --from=builder /app/public /public
COPY .env .env

EXPOSE 8200/tcp

CMD ["./main"]
