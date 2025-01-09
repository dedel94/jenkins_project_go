FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build current_time.go
CMD ["./current_time"]
