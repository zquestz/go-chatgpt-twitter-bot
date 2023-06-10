# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.19.1

LABEL org.opencontainers.image.authors="Josh Ellithorpe <quest@mac.com>"

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/zquestz/go-chatgpt-twitter-bot

# Switch to the correct working directory.
WORKDIR /go/src/github.com/zquestz/go-chatgpt-twitter-bot

# Build the code.
RUN make install

# Set the start command.
ENTRYPOINT ["/go/bin/go-chatgpt-twitter-bot"]
