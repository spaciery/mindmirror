# Step 1: Build the Go application
FROM golang:1.20 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o mindmirror .

# Step 2: Run the Go application
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/mindmirror .
EXPOSE 8081
CMD ["./mindmirror"]
