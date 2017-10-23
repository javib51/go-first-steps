FROM golang:latest

ADD main.go /go/src/main.go

RUN go get -u github.com/gorilla/mux && \
	go get -u github.com/go-sql-driver/mysql && \
	go build -o /go/bin/server /go/src/main.go

ENTRYPOINT ["/go/bin/server"]
