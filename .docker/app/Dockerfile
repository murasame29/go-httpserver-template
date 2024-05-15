FROM golang:1.22.2 AS builder
WORKDIR /build

COPY . .

RUN go build -o app ./cmd/app/main.go

FROM ubuntu AS runner
WORKDIR /app

COPY --from=builder /build/app .

ENV HOST=
ENV PORT=
ENV SHUTDOWN_TIMEOUT=

EXPOSE ${PORT}

CMD ["./app"]