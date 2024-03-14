# Build app
FROM ubuntu:20.04 AS sekai-builder

# Avoid prompts from apt
ARG DEBIAN_FRONTEND=noninteractive

# Update packages and Install essentials in one layer
RUN apt-get update -y && \
    apt-get install -y git build-essential jq wget

# Install Golang
RUN wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz -O /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

# Set PATH for Golang
ENV PATH="$PATH:/usr/local/go/bin"

WORKDIR /app

COPY . /app

RUN make build-static

# Run app
FROM ubuntu:20.04

# Avoid prompts from apt
ARG DEBIAN_FRONTEND=noninteractive

# Copy artifacts
COPY --from=sekai-builder /sekaid /bin/sekaid

# Expose the default Tendermint port
EXPOSE 26657
EXPOSE 26656
EXPOSE 8080
