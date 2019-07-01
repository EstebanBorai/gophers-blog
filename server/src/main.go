package main

import (
	r "github.com/estebanborai/go-server-sample/server/src/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r.StartServer()
}
