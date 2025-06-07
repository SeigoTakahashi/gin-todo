package main

import (
	"gin-todo/infra"
	"gin-todo/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// マイグレーションの実行
	if err := db.AutoMigrate(
		&models.Todo{},
	); err != nil {
		panic("failed to migrate database")
	}
}
