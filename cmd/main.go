package main

import (
	"github.com/gin-gonic/gin"
	"go_book_api/api"
)

func main() {
	api.InitDB()
	r := gin.Default()

	// Public routes
	r.POST("/token", api.GenerateJWT)

	// protected routes
	protected := r.Group("/", api.JWTAuthMiddleware())
	{
		protected.POST("/book", api.CreateBook)
		protected.GET("/books", api.GetBooks)
		protected.GET("/book/:id", api.GetBook)
		protected.PUT("/book/:id", api.UpdateBook)
		protected.DELETE("/book/:id", api.DeleteBook)
	}

	r.Run(":8080")
}
