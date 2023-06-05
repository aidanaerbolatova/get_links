FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base && go build -o app cmd/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/config /app/config
CMD ["./app"]