FROM golang:1.23.6-alpine

RUN apk add --no-cache make git

WORKDIR /build

RUN git clone https://github.com/datvo2k/MOTD.git

RUN go mod download
RUN chmod +x /build/motd.sh
RUN /bin/sh /build/motd.sh

CMD ["/bin/sh", "-c", "bash"]