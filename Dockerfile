## First stage - Build stage
# Pull golang image from Dockerhub
FROM golang:alpine AS builder

# Set up ENV vars
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set up the working directory
WORKDIR /app1

# copy the source code, then run build command
COPY ./modules .
RUN go build -o weather .

## Second stage - Run stage
FROM alpine:latest

# Set up the working directory
WORKDIR /app2

# Copy the executable binary file, env file and config file from the last stage to the new stage
COPY --from=builder /app1/weather .
COPY config.yaml .

ARG PORT
EXPOSE $PORT

# Execute the build
CMD ["./weather"]