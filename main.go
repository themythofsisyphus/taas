package main

import (
	"taas/db"
	"taas/handlers"
	"taas/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	handlers.RegisterTagRoutes(r, db.Connect())

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
