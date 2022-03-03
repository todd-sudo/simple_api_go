package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/todd-sudo/todo/internal/middleware"
	"github.com/todd-sudo/todo/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", h.Login)
		authRoutes.POST("/register", h.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(h.service.JWT))
	{
		userRoutes.GET("/profile", h.ProfileUser)
		userRoutes.PUT("/profile", h.UpdateUser)
	}

	bookRoutes := r.Group("api/items", middleware.AuthorizeJWT(h.service.JWT))
	{
		bookRoutes.GET("/", h.AllItem)
		bookRoutes.POST("/", h.InsertItem)
		bookRoutes.GET("/:id", h.FindByIDItem)
		bookRoutes.PUT("/:id", h.UpdateItem)
		bookRoutes.DELETE("/:id", h.DeleteItem)
	}

	return r
}
