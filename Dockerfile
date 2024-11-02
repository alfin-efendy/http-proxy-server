##
## Build
##
# Start from the latest golang base image
FROM golang:alpine AS build

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/packages/goginapp/

# Fetch dependencies.
# Using go get.
COPY . .
RUN go get -d -v

ENV GO111MODULE="on" \
  GOARCH="amd64" \
  GOOS="linux" \
  CGO_ENABLED="0"

# Build the Go app
RUN go build -o /go/server

############################
# STEP 2 build a small image
############################
# Image for application
FROM alpine:3.16.8

# Copy executable.
COPY --from=build /go/server /app/server

RUN chmod -R 777 /app/server

WORKDIR /app

ENTRYPOINT ["/app/server"]