package controllers

import (
  "fmt"
  "github.com/gin-gonic/gin"
  data "github.com/estebanborai/songs-share-server/data"
  models "github.com/estebanborai/songs-share-server/models"
  eh "github.com/estebanborai/songs-share-server/lib/error_handlers"
)

func FindUserById(c *gin.Context, id string) {
  var user models.User

  db := data.Connection(c)

  if userResult := db.Where(&models.User{ Id: id }).First(&user); userResult.Error == nil {
    c.JSON(200, user)
  } else {
    eh.NotFound(c, fmt.Sprintf("%s doesn't exists", id))
  }
}
