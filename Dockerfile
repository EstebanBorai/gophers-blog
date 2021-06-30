FROM golang:1.16

WORKDIR /go/src/github.com/EstebanBorai/gophers-blog

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

RUN ["go", "mod", "tidy"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o main ./cmd" -command="./main"
