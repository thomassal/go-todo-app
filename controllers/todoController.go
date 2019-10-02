package controllers

import (
	"encoding/json"
	"go-todo-app/models"
	u "go-todo-app/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateTodo = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	todo.UserId = user
	resp := todo.Create()
	u.Respond(w, resp)
}

var RetrieveTodo = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetTodo(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateTodo = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	todo := &models.Todo{}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	error := json.NewDecoder(r.Body).Decode(todo)
	if error != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	todo.UserId = user
	resp := todo.Update(uint(id))
	u.Respond(w, resp)
}

var DeleteTodo = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	resp := models.Delete(uint(id))
	u.Respond(w, resp)
}

var GetTodos = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetTodos(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
