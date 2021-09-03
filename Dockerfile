from golang:1.17.0-alpine3.14

RUN apk add bash

WORKDIR /go/src/app

# RUN go get "github.com/julienschmidt/httprouter"
# RUN go get "gopkg.in/mgo.v2"
# RUN go get "gopkg.in/mgo.v2/bson"

CMD tail -f /dev/null