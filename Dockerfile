# Start from golang base image
FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN export GIN_MODE=release
RUN export SCOPE=prod

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build src/cmd/main.go

# Expose port 50052 to the outside world
EXPOSE 50052

# Run the executable
CMD [ "/build" ]
