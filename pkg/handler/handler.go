package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/poroshin-ao/test-project/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sing-in", h.signIn)
		auth.POST("sing-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)

			items := lists.Group("id/items")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.UpdateItem)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.CreateItem)
			items.DELETE("/:id", h.DeleteItem)
		}

	}

	return router
}
