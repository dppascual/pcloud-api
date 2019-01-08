#############
#
# First stage
#
#############
FROM golang:1.11.2-alpine3.8 as build-env

# Add git
RUN apk update && \
    apk upgrade && \
    apk add git

# Create a workspace
RUN mkdir /api
WORKDIR /api
COPY go.mod .
COPY go.sum .

# Get dependancies - will be cached if mod/sum files won't be changed
RUN go mod download

# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/api

#############
#
# Second stage
#
#############
FROM ubuntu:latest
COPY --from=build-env /go/bin/api /go/bin/api
ENTRYPOINT ["/go/bin/api"]