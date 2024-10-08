# Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/g3 ./cmd/api/main.go

# Run Stage
FROM alpine:latest

WORKDIR /app

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*

COPY --from=builder /build/bin/g3 ./g3

COPY .env /app

COPY --from=builder /build/db /app/db

EXPOSE 3000

CMD [ "/app/g3" ]