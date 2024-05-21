package handler

import (
	"Stepuha.net/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func (handl *Handler) RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handl.signUp)
		auth.POST("/sign-in", handl.signIn)
	}

	api := router.Group("/api", handl.userIdentity)
	{
		userGoods := api.Group("/goods")
		{
			userGoods.POST("/", handl.createGood)
			userGoods.GET("/", handl.getAllGoods)
			userGoods.GET("/:id", handl.getGoodById)
			userGoods.DELETE("/:id", handl.deleteGood)
			userGoods.PUT("/:id", handl.updateGood)
			picture := userGoods.Group("/picture")
			{
				picture.PUT("/:id", handl.uploadPicture)
			}
		}
		users := api.Group("/users")
		{
			users.GET("/:id", handl.getUserById)
			users.GET("/", handl.getYourUser)
			users.PUT("/", handl.updateUser)
		}
	}

	suppApi := api.Group("/supp")
	{
		suppApi.GET("/rnd", handl.getRandomGoods)
	}
}
