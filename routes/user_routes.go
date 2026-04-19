package routes

import (
	"encoding/json"
	"net/http"

	"strconv"
	"user-api/domain/services"
	"user-api/exceptions"
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
	router.Get("/", handler.GetPage)
	router.Get("/{id}", handler.GetByID)
	router.Post("/", handler.Create)
	router.Delete("/{id}", handler.Delete)

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
		utils.HandleError(responseWriter, exceptions.InvalidArgument("invalid id format"))
		return
	}

	user, err := handler.userService.GetByID(id)
	if err != nil {
		utils.HandleError(responseWriter, err)
		return
	}

	utils.OkResponse(responseWriter, mapping.MapToResponse(user))
}

// @Summary      Get paged users
// @Description  Returns a paged list of users
// @Tags         Users
// @Produce      json
// @Param        page   query  int  false  "Page number"   default(1)
// @Param        size   query  int  false  "Page size"     default(10)
// @Success      200    {object}  contracts.PagedResponse[contracts.UserResponse]
// @Failure      500    {string}  string  "internal error"
// @Router       /users [get]
func (handler UserRouteHandler) GetPage(responseWriter http.ResponseWriter, request *http.Request) {
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(request.URL.Query().Get("size"))
	if err != nil || size < 1 {
		size = 10
	}

	users, total, err := handler.userService.GetPage(page, size)
	if err != nil {
		utils.HandleError(responseWriter, exceptions.Internal())
		return
	}

	pagedResponse := mapping.MapToPagedResponse(users, page, size, total)

	utils.OkResponse(responseWriter, pagedResponse)
}

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
		utils.HandleError(responseWriter, exceptions.InvalidArgument("invalid request body"))
		return
	}

	domainCreateRequest := mapping.MapToDomain(createRequest)
	created, err := handler.userService.Create(domainCreateRequest)
	if err != nil {
		utils.HandleError(responseWriter, exceptions.Internal())
		return
	}

	userResponse := mapping.MapToResponse(created)

	utils.CreatedResponse(responseWriter, userResponse)
}

// @Summary      Delete user by ID
// @Description  Deletes a single user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      204  
// @Failure      400  {string}  string  "invalid id"
// @Failure      404  {string}  string  "not found"
// @Router       /users/{id} [delete]
func (handler UserRouteHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.HandleError(responseWriter, exceptions.InvalidArgument("invalid id format"))
		return
	}

	deleteErr := handler.userService.Delete(id)
	if deleteErr != nil {
		utils.HandleError(responseWriter, err)
		return
	}

	utils.NoContentResponse(responseWriter)
}