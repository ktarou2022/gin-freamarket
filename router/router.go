package router

import (
	"gin-fleamarket/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(ic controllers.IItemController) *gin.Engine {
  g := gin.Default()

  g.GET("/sample", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  g.GET("/items", ic.FindAll)
  g.GET("/items/:id", ic.FindById)
  g.POST("/items", ic.Create)
  g.PUT("/items/:id", ic.Update)

  return g
}
