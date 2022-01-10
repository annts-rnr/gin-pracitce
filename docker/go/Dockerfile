FROM golang:1.17.6-alpine

RUN go get -u github.com/cosmtrek/air

WORKDIR /go/src
CMD ["air", "-c", ".air.toml"]
