FROM golang:1.18-alpine AS build

# The current folder is set to app
WORKDIR /app

# First copy the mod file and download the dependencies
# The idea is to seperate the layers so builds without depenceny 
# change will be slightly faster.
COPY go.mod .
RUN go mod download

# Copy the rest of the code
COPY . ./

# Generate a static binary (all in one without external dependencies)
ENV CGO_ENABLED=0
RUN go build -o restapp ./cmd/main.go

FROM alpine:3.18.4

# Create a non-root user to run the app
RUN adduser -D -u 8118 appuser
USER appuser:appuser

# Set the current working directory
WORKDIR /app


# Copy the builded file
COPY --from=build --chown=appuser:appuser /app/restapp .

# Before running expose the port
EXPOSE 80

# Run the app
CMD ["./restapp"]
