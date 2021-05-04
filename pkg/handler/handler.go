package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	auth := r.Group("/auth")
	{
		auth.POST("/sing-up")
		auth.POST("/sign-in")
	}
	api := r.Group("/api/v1")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")
			items := lists.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				lists.PUT("/:item_id")
				lists.DELETE("/:item_id")
			}
		}
	}
	return r
}
