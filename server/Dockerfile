FROM golang:1.12

COPY ./src/ $GOPATH/src/github.com/estebanborai/go-server-sample/server/src/

WORKDIR $GOPATH/src/github.com/estebanborai/go-server-sample/server/src/

RUN go get -u -v github.com/golang/dep/...
RUN dep ensure

EXPOSE 8080

CMD ["go", "run", "main.go"]
