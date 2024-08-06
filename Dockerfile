FROM golang:1.19 AS builder

WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o makvesapp -ldflags="-s -w" ./cmd/app/main.go

FROM alpine:latest

RUN apk update && apk add curl

WORKDIR /

COPY --from=builder /app/makvesapp /
COPY --from=builder /app/config /config

EXPOSE 8084

ENTRYPOINT ["/makvesapp"]