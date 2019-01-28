#################################
# STEP 1 setup env
#################################
FROM alpine as builder

# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser


#############################
# STEP 2 build a small image
#############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /go/src/app
COPY  bin/animal-api animal-api

# Use an unprivileged user.
USER appuser

EXPOSE 8080

# Run the binary
CMD ["./animal-api"]
