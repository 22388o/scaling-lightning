# Run from the root of scaling lightning
# docker build -t user/bitcoind -f clients/bitcoind/Dockerfile .

# build stage
FROM golang:buster as backend-builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY pkg/ ./pkg
COPY clients/bitcoind/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o bitcoind .

# final stage
FROM debian:buster-slim
COPY --from=backend-builder /app/bitcoind /app/
RUN useradd -ms /bin/bash bitcoind
RUN apt-get -y update && apt-get -y --no-install-recommends install ca-certificates bash && rm -rf /var/lib/apt/lists/*;
RUN update-ca-certificates
WORKDIR /app
USER bitcoind
ENTRYPOINT ["./bitcoind"]