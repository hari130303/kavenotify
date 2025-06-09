# syntax=docker/dockerfile:1

# Stage 1: Prepare CA certificates
FROM alpine:latest AS certs

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Stage 2: Create a minimal image with the compiled binary
FROM scratch

# Copy the prebuilt binary into the final image
COPY kavenotify-binary /kavenotify-binary

# Copy the CA certificates from the certs stage
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

# Run the binary
CMD ["/kavenotify-binary"]
