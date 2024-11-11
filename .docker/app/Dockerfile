FROM golang:1.23.2 AS builder
WORKDIR /build

COPY . .

RUN CGO_ENABLED=0  go build -o /build/main ./cmd/app/main.go

FROM gcr.io/distroless/static-debian11:nonroot
WORKDIR /app

COPY --from=builder /build/main /

ENV PORT=

EXPOSE ${PORT}

USER 1000

CMD ["/main"]