package router

import (
	"gin-fleamarket/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(ic controllers.IItemController, ac controllers.IAuthController) *gin.Engine {
  g := gin.Default()
  itemRouter := g.Group("/items")
  itemRouter.GET("", ic.FindAll)
  itemRouter.GET("/:id", ic.FindById)
  itemRouter.POST("", ic.Create)
  itemRouter.PUT("/:id", ic.Update)
  itemRouter.DELETE("/:id", ic.Delete)

  authRouter := g.Group("/auth")
  authRouter.POST("/signup", ac.Siginup)
  authRouter.POST("/login", ac.Login)

  return g
}
