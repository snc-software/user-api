package routes

import (
	"encoding/json"
	"net/http"

	//"strconv"
	domainerrors "user-api/domain/errors"
	"user-api/domain/services"
	"user-api/mapping"
	"user-api/routes/contracts"
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
	//router.Get("/", handler.GetPage)
	router.Get("/{id}", handler.GetByID)
	router.Post("/", handler.Create)

	return router
}

// @Summary      Get user by ID
// @Description  Returns a single user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  contracts.UserResponse
// @Failure      400  {string}  string  "invalid id"
// @Failure      404  {string}  string  "not found"
// @Router       /users/{id} [get]
func (handler UserRouteHandler) GetByID(responseWriter http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.Error(responseWriter, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := handler.userService.GetByID(id)
	if err == domainerrors.NotFound {
		utils.Error(responseWriter, "user not found", http.StatusNotFound)
		return
	}
	if err != nil {
		utils.Error(responseWriter, "internal error", http.StatusInternalServerError)
		return
	}

	userResponse := mapping.MapToResponse(user)

	utils.JSON(responseWriter, userResponse, http.StatusOK)
}

// // @Summary      Get paged users
// // @Description  Returns a paged list of users
// // @Tags         Users
// // @Produce      json
// // @Param        page   query  int  false  "Page number"   default(1)
// // @Param        size   query  int  false  "Page size"     default(10)
// // @Success      200    {object}  models.PagedResponse[models.UserModel]
// // @Failure      500    {string}  string  "internal error"
// // @Router       /users [get]
// func (handler UserRouteHandler) GetPage(responseWriter http.ResponseWriter, request *http.Request) {
// 	page, err := strconv.Atoi(request.URL.Query().Get("page"))
// 	if err != nil || page < 1 {
// 		page = 1
// 	}

// 	size, err := strconv.Atoi(request.URL.Query().Get("size"))
// 	if err != nil || size < 1 {
// 		size = 10
// 	}

// 	result, err := handler.userService.GetPage(page, size)
// 	if err != nil {
// 		utils.Error(responseWriter, "internal error", http.StatusInternalServerError)
// 		return
// 	}

// 	utils.JSON(responseWriter, result, http.StatusOK)
// }

// @Summary      Create user
// @Description  Creates a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      contracts.CreateUserRequest  true  "User"
// @Success      201   {object}  contracts.UserResponse
// @Failure      400   {string}  string  "invalid request body"
// @Failure      500   {string}  string  "internal error"
// @Router       /users [post]
func (handler UserRouteHandler) Create(responseWriter http.ResponseWriter, request *http.Request) {
	var createRequest contracts.CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&createRequest)
	if err != nil {
		utils.Error(responseWriter, "invalid request body", http.StatusBadRequest)
		return
	}

	domainCreateRequest := mapping.MapToDomain(createRequest)
	created, err := handler.userService.Create(domainCreateRequest)
	if err != nil {
		utils.Error(responseWriter, "internal error", http.StatusInternalServerError)
		return
	}

	userResponse := mapping.MapToResponse(created)

	utils.JSON(responseWriter, userResponse, http.StatusCreated)
}