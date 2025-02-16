FROM golang:1.23.6

RUN apt-get update && apt-get install -y \
    gcc-multilib \
    g++-multilib \
    crossbuild-essential-armhf \
    crossbuild-essential-arm64 \
    && rm -rf /var/lib/apt/lists/*

ENV GOOS=linux
ENV GOARCH=amd64