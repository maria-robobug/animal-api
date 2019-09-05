#################################
# STEP 1 setup env
#################################
FROM golang:1.12.1-alpine3.9 as builder

RUN apk add --update --no-cache git ca-certificates

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /app

COPY . .

ARG VERSION_ARG=0.0.0

RUN go mod download
RUN go build -o animal-api-${VERSION_ARG}

#############################
# STEP 2 build a small image
#############################
FROM scratch

ARG VERSION_ARG=0.0.0
ENV VERSION=${VERSION_ARG}

EXPOSE 8080

# Import files from the builder.
COPY --from=builder /app/animal-api-${VERSION} /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT /app/animal-api-${VERSION}
