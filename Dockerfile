FROM golang:1.16

WORKDIR /go/src/github.com/EstebanBorai/gosk

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o main ./cmd" -command="./main"
