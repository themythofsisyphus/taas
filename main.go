package main

import (
	"log"
	"net/http"
	"taas/config"
	"taas/db"
	"taas/middleware"
	"taas/repository"
	"taas/router"
	"taas/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// middlewares
	r.Use(middleware.AuthMiddleware())
	r.Use(middleware.Logger())

	// config
	configs, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Configs couldn't loaded")
	}
	db := db.InitDB(&configs.Database)

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos)
	router.RegisterRoutes(r, services)

	// http server
	srv := &http.Server{
		Addr:         ":" + configs.Server.Port,
		Handler:      r,
		ReadTimeout:  configs.Server.ReadTimeOut,
		WriteTimeout: configs.Server.WriteTimeOut,
		IdleTimeout:  configs.Server.IdleTimeOut,
	}

	log.Printf("Server starting on port %s", configs.Server.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server failed to start:", err)
	}
}
