FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ya-storage ./cmd/ya-storage

FROM alpine

WORKDIR /app

COPY --from=builder /build/ya-storage /app/ya-storage
COPY --from=builder /build/config /app/config

EXPOSE 8086

ENTRYPOINT ["/app/ya-storage"]