FROM golang:1.7-wheezy

WORKDIR /go/src/github.com/waltzofpearls/sawmill

COPY . .
RUN make

COPY docker/config.docker.yml config.yml

EXPOSE 9000

CMD ["./sawmill"]
