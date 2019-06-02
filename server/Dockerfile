FROM golang:1.12

COPY ./src/ $GOPATH/src/github.com/estebanborai/songs-share-server/

WORKDIR $GOPATH/src/github.com/estebanborai/songs-share-server/


RUN go get -u -v github.com/golang/dep/...
# RUN cd $GOPATH/src/github.com/estebanborai/songs-share-server/ && dep init
RUN dep ensure

EXPOSE 8080
