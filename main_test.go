package main

import (
	"encoding/json"
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/router"
	"gin-fleamarket/services"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 他のテスト実行前に呼び出される
func TestMain(m *testing.M) {

	code := m.Run()

	os.Exit(code)
}

func setupTestData() []models.Item {
	items := []models.Item{
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, ID: 1},
		{Name: "テストアイテム2", Price: 2000, Description: "テスト2", SoldOut: true, ID: 2},
		{Name: "テストアイテム3", Price: 3000, Description: "テスト3", SoldOut: false, ID: 3},
	}

	return items
}

func SetUp() *gin.Engine {

	items := setupTestData()
	itemRepository := repositories.NewItemMemoryRepository(items)
	itemServices := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemServices)

	return router.NewRouter(itemController)
}

func TestFindAll(t *testing.T) {
	router := SetUp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}