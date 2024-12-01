# Stage 1: Build the Go application
FROM golang:1.22.2-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

# Stage 2: Create the final image
FROM alpine:latest

COPY --from=builder /app/main /main
COPY --from=builder /app/wait-for-it.sh /wait-for-it.sh
COPY --from=builder /app/configs/config.env /configs/config.env

RUN chmod +x /wait-for-it.sh

EXPOSE 8000

CMD ["/wait-for-it.sh", "db", "5432", "--", "/main"]
