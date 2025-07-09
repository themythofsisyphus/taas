package main

import (
	"taas/db"
	"taas/handler"
	"taas/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	db := db.Connect()

	handler.RegisterTagRoutes(r, db)
	handler.RegisterEntityRoutes(r, db)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
