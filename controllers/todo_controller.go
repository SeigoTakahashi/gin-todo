package controllers

import (
	"gin-todo/dto"
	"gin-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// タスクの作成
func CreateTodo(c *gin.Context, db *gorm.DB) {

	// 入力チェック
	var input dto.CreateTodoInput
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "todo/todo.html", gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Task:      input.Task,
		Completed: false,
	}

	// データベースに作成
	if err := db.Create(&todo).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "todo/todo.html", gin.H{"error": "保存に失敗しました"})
		return
	}

	// 一覧画面にリダイレクト
	c.Redirect(http.StatusSeeOther, "/todo")
}

// タスクの全件取得
func ShowTodo(c *gin.Context, db *gorm.DB) {
	var todos []models.Todo
	// データベースから全件取得
	if err := db.Order("created_at desc").Find(&todos).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "todo/todo.html", gin.H{"error": "データ取得に失敗しました"})
		return
	}

	// 取得したデータをテンプレートに渡して表示
	c.HTML(http.StatusOK, "todo/todo.html", gin.H{
		"title": "ToDoApp",
		"todos": todos,
	})
}

// タスクの完了状態を更新
func CompleteTodoHandler(c *gin.Context, db *gorm.DB) {
	// URLパラメータの取得
	id := c.Param("id")

	var todo models.Todo

	// 対象のタスクを取得
	if err := db.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// 入力チェック
	var request struct {
		Completed bool `json:"completed"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 完了状態を更新
	todo.Completed = request.Completed
	if err := db.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update todo"})
		return
	}

	// 結果をJSONで返す
	c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
}

// タスクの削除
func DeleteTodo(c *gin.Context, db *gorm.DB) {
	// URLパラメータからIDを取得
	id := c.Param("id")

	// 対象のタスクの削除
	if err := db.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to delete the task",
			"deleted": false,
		})
		return
	}

	// 結果をJSONで返す
	c.JSON(http.StatusOK, gin.H{
		"message": "task successfully deleted",
		"deleted": true,
	})
}
