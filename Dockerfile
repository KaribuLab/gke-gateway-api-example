FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go mod download && go build -o /app/main

FROM alpine:3.20

COPY --from=builder /app/main /app/main

CMD ["/app/main"]