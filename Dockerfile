FROM golang:1.22-alpine AS builder

RUN go version

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bot-app ./cmd/main.go



FROM alpine:latest AS final

WORKDIR /app

RUN mkdir -p /app/tmp/video /app/tmp/audio
RUN apk add --no-cache python3 ffmpeg yt-dlp

COPY --from=builder /app/bot-app .
COPY --from=builder /app/internal ./internal

CMD ["./bot-app"]
