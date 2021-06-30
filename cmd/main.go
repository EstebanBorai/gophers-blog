package main

import "github.com/EstebanBorai/gophers-blog/pkg"

func main() {
	repository := pkg.NewRepository()
	server := pkg.NewServer(&repository)

	server.Serve()
}
