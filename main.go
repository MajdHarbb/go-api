package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/auth/signup", signup)
	router.POST("/auth/login", login)

	router.GET("/block/delete/:id", blockDelete)
	router.POST("/block/create", blockCreate)
	router.GET("/block/read/:id", blockReadByUserID)

	router.GET("/like/delete/:id", likeDelete)
	router.POST("/like/create", likeCreate)
	router.GET("/like/read/:id", likeReadByUserID)

	router.Run("localhost:8080")
}
