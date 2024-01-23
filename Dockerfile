FROM golang:1.21.4-alpine AS builder
WORKDIR /build

COPY . .

RUN go build -o app .

FROM ubuntu AS runner
WORKDIR /app

COPY --from=builder /build/app .

ENV HOST=
ENV PORT=
ENV SHUTDOWN_TIMEOUT=

ENV NEW_RELIC_APP_NAME=
ENV NEW_RELIC_LICENSE=

EXPOSE ${PORT}

CMD ["./app"]