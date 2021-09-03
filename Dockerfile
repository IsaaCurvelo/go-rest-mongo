from golang:1.17.0-alpine3.14

RUN apk add bash

WORKDIR /go/src/app

# RUN go get

CMD tail -f /dev/null