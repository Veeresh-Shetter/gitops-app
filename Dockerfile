# --- Stage 1: Build the hardened static binary ---
FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod main.go ./

# Strip debug symbols (-s -w) for a smaller attack surface and binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# --- Stage 2: Create minimal non-root runtime image ---
FROM alpine:latest

# Create a dedicated non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/main .
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

EXPOSE 8080

CMD ["./main"]
