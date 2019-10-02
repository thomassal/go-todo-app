package main

import (
	"fmt"
	"go-todo-app/app"
	"go-todo-app/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/todo/new", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todo/{id}", controllers.RetrieveTodo).Methods("GET")
	router.HandleFunc("/api/todo/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todo/{id}", controllers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET") //  get todos for the user

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
