FROM golang:alpine AS builder

# ENVs
ENV CGO_ENABLED 0 \
    GOPATH /go \
    GOCACHE /go-build \
    GOOS linux \
    GOARCH=amd64 \
    GO111MODULE=on

# Copy
WORKDIR /go-auth
COPY . /go-auth

# Install Dependencies
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go mod download

# Build the binary
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go build -o main main.go


# Start a new stage from scratch
FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /usr/local/bin/app
COPY --from=builder /go-auth/.env .
COPY --from=builder /go-auth/main .
CMD ["./main", "app", "bootstrap"]
