package helpers

import (
  "github.com/gin-gonic/gin"
)

func ContextRequestBody(c *gin.Context) string {
  buf := make([]byte, 1024)
  num, _ := c.Request.Body.Read(buf)
  return string(buf[0:num])
}
