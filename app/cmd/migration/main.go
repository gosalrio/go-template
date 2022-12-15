package main

import (
	// todo "go-todo-app/src/todo/data"
	// user "go-todo-app/src/user/data"
	"go-todo-app/src/utils"
)

func main() {
	utils.LoadEnv()
	db := utils.ShouldInitializeDB()
	db.AutoMigrate(
	// todo.LocalTodo{},
	// user.LocalUser{},
	// user.LocalUserToken{},
	)
}
