FROM golang:1.17-alpine AS builder
LABEL stage=builder
WORKDIR /app
COPY . .

RUN apk add build-base && go build -o test_project cmd/main.go

FROM alpine:3.6
WORKDIR /app
LABEL authors="AidanaErbolatova" project="test_project"
COPY --from=builder /app .
CMD ["/app/test_project"]