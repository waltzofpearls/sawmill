FROM golang:1.7-wheezy

WORKDIR /go/src/app

COPY . .
RUN make all

CMD ["./sawmill"]
