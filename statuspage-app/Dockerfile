# ==========================================
# STAGE 1: Build the Go Binary
# ==========================================
FROM golang:1.26.3-alpine AS builder

# Install git if needed (required if pulling private modules, otherwise optional)
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker layer caching
COPY go.mod ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Accept the build-time arguments passed from the Makefile
ARG GIT_SHA="unknown"
ARG BUILD_TIME="unknown"

# Compile the binary staticly. 
# We inject the variables into the main package via -ldflags.
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X main.CommitSHA=${GIT_SHA} -X main.buildDate=${BUILD_TIME}" \
    -o statuspage-app ./...

# ==========================================
# STAGE 2: Secure, Lightweight Runtime
# ==========================================
FROM gcr.io/distroless/static-debian12:latest

# Set working directory for clarity
WORKDIR /

# Copy ONLY the compiled binary from the builder stage
COPY --from=builder /app/statuspage-app /statuspage-app

# Expose the application port (adjust if your app uses a different port)
EXPOSE 8080

# Run the binary directly (distroless has no shell, so use exec form)
ENTRYPOINT ["/statuspage-app"]