FROM golang:1.22

WORKDIR /app

# Download Go modules
COPY ./src ./

RUN go mod download

# Set destination for COPY

RUN mkdir bin

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["bin/main"]
