package data

import (
	eh "github.com/estebanborai/songs-share-server/server/src/lib/error_handlers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var connectionString string = "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local"

func Connection(c *gin.Context) (db *gorm.DB) {
	db, err := gorm.Open("mysql", connectionString)

	if err != nil {
		eh.ResponseWithError(c, 500, "Unable to connect to the database")
	}

	return db
}
