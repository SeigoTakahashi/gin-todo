package dto

// DTOの定義
type CreateTodoInput struct {
	Task string `form:"task" binding:"required"`
}
