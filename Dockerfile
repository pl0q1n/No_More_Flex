FROM golang:latest

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build 

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./No_More_Flex"]