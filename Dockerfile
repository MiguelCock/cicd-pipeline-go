FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go .
COPY ./endpoints/ ./endpoints/
COPY ./utils/ ./utils/
COPY templates ./templates

RUN go build -o server .

FROM debian:bookworm-slim

RUN useradd -m appuser

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates

RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 5000

CMD ["./server"]