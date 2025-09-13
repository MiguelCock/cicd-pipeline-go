FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates

EXPOSE 5000

CMD ["./server"]