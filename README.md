## go-todo-app

An example API written in go.

### Prerequisites

* GO
* Postgresql

### Packages

* gorilla/mux
* jinzhu/gorm
* dgrijalva/jwt-go
* joho/godotenv

### API Endpoints

#### Users

* **/api/user/new** (User registration endpoint)
* **/api/user/login** (User login endpoint)


#### Todos

* **/api/todos** (Todo list endpoint)
* **/api/todo/new** (Todo create endpoint)
* **/api/todo/{todo-id}** (Todo retrieve, update and destroy endpoint)

### RUN

    go run main.go
