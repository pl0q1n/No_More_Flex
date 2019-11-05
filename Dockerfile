FROM golang:latest AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build ./cmd/nmf-server

FROM alpine:latest
COPY --from=builder /app/nmf-server ./
EXPOSE 8080
ENV PORT 8080
ENV HOST 0.0.0.0
CMD ["./nmf-server"]
