package main

import (
	"github.com/gin-gonic/gin"
	"taas/db"
	"taas/handlers"
)

func main() {
	r := gin.Default()

	handlers.RegisterTagRoutes(r, db.Connect())

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
