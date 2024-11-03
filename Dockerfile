FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/umbrela-test-task

RUN CGO_ENABLED=0 GOOS=linux go build -o echo-hello .

FROM alpine:latest

COPY --from=builder /app/cmd/umbrela-test-task/echo-hello .

CMD ["./echo-hello"]
