package routes

import (
	"encoding/json"
	"net/http"
	"user-api/domain/models"
	"user-api/domain/services"
	"user-api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserRouteHandler struct {
	userService services.UserService
}

func UserRoutes(userService services.UserService) http.Handler {
	handler := UserRouteHandler{userService: userService}

	router := chi.NewRouter()
	router.Get("/{id}", handler.GetByID)
	router.Post("/", handler.Create)

	return router
}

// @Summary      Get user by ID
// @Description  Returns a single user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.UserModel
// @Failure      400  {string}  string  "invalid id"
// @Router       /users/{id} [get]
func (handler UserRouteHandler) GetByID(responseWriter http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(responseWriter, "invalid id", http.StatusBadRequest)
		return
	}
	user := handler.userService.GetByID(id)
	utils.JSON(responseWriter, user, http.StatusOK)
}

// @Summary      Create user
// @Description  Creates a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      models.CreateUserModel  true  "User"
// @Success      201   {object}  models.UserModel
// @Failure      400   {string}  string  "invalid request body"
// @Router       /users [post]
func (handler UserRouteHandler) Create(responseWriter http.ResponseWriter, request *http.Request) {
	var createRequest models.CreateUserModel
	err := json.NewDecoder(request.Body).Decode(&createRequest)
	if err != nil {
		http.Error(responseWriter, "invalid request body", http.StatusBadRequest)
		return
	}
	created := handler.userService.Create(createRequest)
	utils.JSON(responseWriter, created, http.StatusCreated)
}