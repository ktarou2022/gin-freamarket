package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	// "gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/router"
	"gin-fleamarket/services"
)

func main() {
  infra.Initialize()
  db := infra.SetupDB()
  // items := []models.Item{
  //   {ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
  //   {ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
  //   {ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
  // }

  // itemRepository := repositories.NewItemMemoryRepository(items)
  itemRepository := repositories.NewItemRepository(db)
  itemServices := services.NewItemService(itemRepository)
  itemController := controllers.NewItemController(itemServices)

  authRepository := repositories.NewAuthRepository(db)
  authSevice := services.NewAuthService(authRepository)
  authController := controllers.NewAuthController(authSevice)
  g := router.NewRouter(itemController, authController)

  g.Run(":8080") // デフォルトで0.0.0.0:8080で待機します
}