package handlers

import (
	"errors"
	"musicapi/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MapRoutes Создает маршруты
func MapRoutes(router *gin.Engine) error {
	if router == nil {
		return errors.New("o my god, router not created")
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is simple music api from Mukhamedjanov Erjan")
	})

	musicRouterGroup := router.Group("/music")
	{
		musicRouterGroup.GET("", GetMusics)
		musicRouterGroup.GET("/:id", GetMusic)
		musicRouterGroup.Use(middlewares.CheckToken)
		musicRouterGroup.POST("", AddMusic)
		musicRouterGroup.DELETE("/:id", DeleteMusic)
	}
	artistRouterGroup := router.Group("/artist")
	{
		artistRouterGroup.GET("", GetArtists)
		artistRouterGroup.GET("/:id", GetArtist)
		artistRouterGroup.Use(middlewares.CheckToken)
		artistRouterGroup.POST("", AddArtist)
		artistRouterGroup.DELETE("/:id", DeleteArtist)
	}
	userRouterGroup := router.Group("/user")
	{

		userRouterGroup.POST("/login", Login)
		userRouterGroup.POST("/signup", Signup)
		userRouterGroup.Use(middlewares.CheckToken)
		userRouterGroup.DELETE("", DeleteUser)
	}
	return nil
}
