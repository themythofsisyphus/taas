// Package main is the entry point of the Tags as a Service (TaaS) application.
package main

import (
	"net/http"
	"taas/config"
	"taas/db"
	"taas/middleware"
	"taas/pkg/tlog"
	"taas/repository"
	"taas/router"
	"taas/service"
	"taas/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// config
	configs, err := config.LoadConfig()
	if err != nil {
		tlog.Fatal("Configs couldn't loaded: %v", err)
	}

	db := db.InitDB(&configs.Database)
	cache := utils.NewCacheClient(&configs.Memcache)

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos, cache)

	// middlewares
	r.Use(middleware.AuthMiddleware(&configs.JWTSecret, services.Tenant))
	r.Use(middleware.Logger())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(tlog.TlogMiddleware())  // lets use our own logger for gin - commented out for now

	router.RegisterRoutes(r, services)

	// http server
	srv := &http.Server{
		Addr:         ":" + configs.Server.Port,
		Handler:      r,
		ReadTimeout:  configs.Server.ReadTimeOut,
		WriteTimeout: configs.Server.WriteTimeOut,
		IdleTimeout:  configs.Server.IdleTimeOut,
	}

	tlog.Info("Server starting on port %s", configs.Server.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		tlog.Fatal("Server failed to start: %v", err)
	}
}
