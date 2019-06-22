package data

import (
	eh "github.com/estebanborai/songs-share-server/server/src/lib/error_handlers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var connectionString = "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local"
var localConnectionString = "root:root@tcp(33060)/songs-share?charset=utf8mb4&parserTime=True&loc=Local"

// Connection Creates a connection to the database and returns it
func Connection(c *gin.Context) (db *gorm.DB) {
	db, err := gorm.Open("mysql", localConnectionString)

	if err != nil {
		eh.SpecialError(c, "Unable to connect to the database", err)
	}

	return db
}
