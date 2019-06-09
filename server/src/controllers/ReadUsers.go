package controllers

import (
  "github.com/gin-gonic/gin"
  models "github.com/estebanborai/songs-share-server/models"
  data "github.com/estebanborai/songs-share-server/data"
)

func ReadUsers(c *gin.Context) {
  var users []models.User

  db := data.Connection(c)
  
  db.Find(&users)

  c.JSON(200, users)
}
