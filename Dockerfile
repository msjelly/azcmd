
# Build the azcmd binary
FROM golang:1.10.3 as builder

# Copy in the go src
WORKDIR /go/src/github.com/msjelly/azcmd
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/azcmd github.com/msjelly/azcmd/cmd/azcmd

# Copy the azcmd into a thin image
FROM ubuntu:latest

# https://blog.cloud66.com/x509-error-when-using-https-inside-a-docker-container/
RUN apt-get update \
    && apt-get install -y ca-certificates
WORKDIR /
COPY --from=builder /go/src/github.com/msjelly/azcmd/bin/azcmd .

ENTRYPOINT ["/azcmd"]