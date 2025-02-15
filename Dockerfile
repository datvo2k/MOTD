FROM golang:1.23.6-alpine

RUN apk add --no-cache make

WORKDIR /build
COPY . .
RUN go mod download
RUN chmod +x /build/motd.sh
RUN /bin/sh /build/motd.sh

CMD ["/bin/sh", "-c", "bash"]