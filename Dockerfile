# --- Stage 1: Build the hardened static binary ---
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY main.go .

# Strip debug symbols (-s -w) for a smaller attack surface and binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main main.go

# --- Stage 2: Create minimal non-root runtime image ---
FROM alpine:3.20

# Create a dedicated non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/main .
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

EXPOSE 8080

CMD ["./main"]
