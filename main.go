package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/reminders", func(c *gin.Context) {
		c.Status(200)
	})

	router.Run()
}
