package main

import (
	"gin-todo/controllers"
	"gin-todo/infra"

	"github.com/gin-gonic/gin"
)

func main() {
	// DBの初期化
	infra.Initialize()
	db := infra.SetupDB()

	// デフォルトルート
	router := gin.Default()

	// 静的ファイルの提供（BootstrouterapやCSSなど）
	router.Static("/static", "./static")

	// テンプレートファイルの読み込み
	router.LoadHTMLGlob("templates/**/*")

	// ToDoルート
	todoRouter := router.Group("/todo")
	todoRouter.GET("", func(c *gin.Context) {
		controllers.ShowTodo(c, db)
	})
	todoRouter.POST("/add", func(c *gin.Context) {
		controllers.CreateTodo(c, db)
	})
	todoRouter.PATCH("/complete/:id", func(c *gin.Context) {
		controllers.CompleteTodoHandler(c, db)
	})
	todoRouter.DELETE("/delete/:id", func(c *gin.Context) {
		controllers.DeleteTodo(c, db)
	})

	router.Run(":8080") // localhost:8080 で起動
}
