package main

import (
	"fmt"
	"net/http"
	_ "user-api/docs"
	"user-api/domain/services"
	"user-api/routes"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           User API
// @version         1.0
// @description     A basic user management API
// @host            localhost:8080
// @BasePath        /
// @tag.name        Users
// @tag.description Operations for managing users
func main() {
	godotenv.Load(".env")
	godotenv.Overload(".env.local")
	userService := services.UserService{}
	
	router := chi.NewRouter()
	router.Mount("/users", routes.UserRoutes(userService))
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	fmt.Println("Server running on http://localhost:8080/swagger/index.html")
	http.ListenAndServe(":8080", router)
}