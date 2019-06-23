package main

import (
	r "github.com/estebanborai/songs-share-server/server/src/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r.StartServer()
}
