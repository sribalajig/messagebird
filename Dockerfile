FROM golang:1.9

WORKDIR /go/src/messagebird

COPY . .

RUN go get github.com/messagebird/go-rest-api \
	&& go get -u github.com/gorilla/mux \
	&& go get github.com/satori/go.uuid \
	&& go get gopkg.in/mgo.v2


ENTRYPOINT go run /go/src/messagebird/cmd/api/main.go