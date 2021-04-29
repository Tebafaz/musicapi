package main

import (
	"context"
	"errors"
	"fmt"
	"musicapi/database"
	"musicapi/handlers"
	"musicapi/middlewares"
	"musicapi/redis"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "musicapi/docs"
)

func init() {
	database.InitPostgres()
	redis.InitRedis()
}

// @title Swagger For Music Api
// @version 1.0
// @description By Mukhamedjanov Erjan
// @termsOfService http://swagger.io/terms/

// @contact.name Mukhamedjanov Erjan
// @contact.url http://www.swagger.io/support
// @contact.email muhamedzhanov.erzhan@mail.ru

// @host localhost:8080
// @BasePath /

func startServer() (*http.Server, <-chan error) {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(middlewares.AccessLog(false))
	router.Use(gin.Recovery())

	swaggerURL := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	returnChannel := make(chan error, 1)

	err := handlers.MapRoutes(router)
	if err != nil {
		returnChannel <- err
		return nil, returnChannel
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			returnChannel <- err
		}
	}()
	return srv, returnChannel
}

func main() {

	server, startServerError := startServer()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Server started")
	select {
	case sig := <-quit:
		fmt.Println(fmt.Sprintf("Received signal: %s\nStopping server...", sig.String()))
	case err := <-startServerError:
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	fmt.Println("Server stopped")
}
