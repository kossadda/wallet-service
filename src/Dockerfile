FROM golang:1.22.2-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build cmd/main.go
EXPOSE 8000
CMD ["./main"]
