FROM golang:1.7-wheezy

WORKDIR /go/src/github.com/waltzofpearls/sawmill

COPY . .
RUN make all

COPY docker/config.docker.yml config.yml
CMD ["./sawmill"]
