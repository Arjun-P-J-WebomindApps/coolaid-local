# Stage 1: Build the Go binary
FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

# Stage 2: Use distroless minimal base image
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /go/bin/app /
COPY .env /home/nonroot/.env

# Run the binary as non-root user
USER nonroot

CMD ["/app"]
