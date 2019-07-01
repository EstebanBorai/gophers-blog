package data

import (
	"log"
	"os"

	"github.com/estebanborai/go-server-sample/server/src/helpers/gimlet"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DockerEnv is the connection string for Docker Environment
var DockerEnv = "root:root@tcp(go-server-sample-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local"

// LocalEnv is the connection string for Go Command (when running locally)
// NOTE: docker port <container id> 3306 will gather the IP of your docker container running the database
var LocalEnv = "root:root@tcp(0.0.0.0:3306)/songs-share?charset=utf8mb4&parseTime=True&loc=Local"

// GetConnectionString returns the connection string for the current environment
func GetConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ServerEnv := os.Getenv("SERVER_ENV")

	if ServerEnv == "LOCAL" {
		return LocalEnv
	}

	return DockerEnv
}

// Connection Creates a connection to the database and returns it
func Connection(c *gin.Context) (db *gorm.DB) {
	db, err := gorm.Open("mysql", GetConnectionString())

	if err != nil {
		gimlet.InternalServerError(c, "Unable to connect to the database")
	}

	return db
}
