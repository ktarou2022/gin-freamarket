package router

import (
	"gin-fleamarket/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(ic controllers.IItemController) *gin.Engine {
  g := gin.Default()
  itemRouter := g.Group("/items")
  itemRouter.GET("", ic.FindAll)
  itemRouter.GET("/:id", ic.FindById)
  itemRouter.POST("", ic.Create)
  itemRouter.PUT("/:id", ic.Update)
  itemRouter.DELETE("/:id", ic.Delete)

  return g
}
