FROM golang:latest

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build ./cmd/nmf-server

# Expose port 8080 to the outside world
EXPOSE 8080
ENV PORT 8080
ENV HOST 0.0.0.0

# Command to run the executable
CMD ["./nmf-server"]
