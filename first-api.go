package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hii nandkumar Khomane This Side This is My My rest API Using Go Lang",
			"status":  "200 ok",
		})
	})
	r.GET("/intro", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":       "Hii my Name is Nandkumar Khomane ",
			"education:": "i am currently pursuing master in computer science",
			"work":       "i am currently working as Assosiate Software Developer At a lentra.ai",
			"address":    "i am from Mandavgan Pharata currently stay in pune for job purpose",
		})
	})
	r.Run("localhost:8085") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
