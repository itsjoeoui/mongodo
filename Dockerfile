FROM golang:1.19

# Set up workspace
RUN mkdir /workspace
WORKDIR /workspace
ADD . .

# Build and test
RUN go build
RUN go test
