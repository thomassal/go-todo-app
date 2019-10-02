package models

import (
	"fmt"
	u "go-todo-app/utils"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name   string `json:"name"`
	Done   bool   `json:"done"`
	UserId uint   `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (todo *Todo) Validate() (map[string]interface{}, bool) {

	if todo.Name == "" {
		return u.Message(false, "Todo name should be on the payload"), false
	}

	if todo.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

/*Create a todo*/
func (todo *Todo) Create() map[string]interface{} {

	if resp, ok := todo.Validate(); !ok {
		return resp
	}

	GetDB().Create(todo)

	resp := u.Message(true, "success")
	resp["todo"] = todo
	return resp
}

/*Update a todo*/
func (todo *Todo) Update(id uint) map[string]interface{} {

	if resp, ok := todo.Validate(); !ok {
		return resp
	}
	todo.ID = id
	err := GetDB().Table("todos").Where("id = ?", id).Update(todo).Error
	if err != nil {
		return nil
	}

	resp := u.Message(true, "Updated")
	resp["todo"] = todo
	return resp
}

func Delete(id uint) map[string]interface{} {

	todo := &Todo{}
	err := GetDB().Table("todos").Where("id = ?", id).Delete(todo).Error
	if err != nil {
		return nil
	}

	resp := u.Message(true, "Deleted")
	return resp
}

func GetTodo(id uint) *Todo {

	todo := &Todo{}
	err := GetDB().Table("todos").Where("id = ?", id).First(todo).Error
	if err != nil {
		return nil
	}
	return todo
}

func GetTodos(user uint) []*Todo {

	todos := make([]*Todo, 0)
	err := GetDB().Table("todos").Where("user_id = ?", user).Find(&todos).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return todos
}
