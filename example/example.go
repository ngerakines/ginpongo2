package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngerakines/ginpongo2"
	"log"
)

func main() {
	log.Println("starting")
	r := gin.Default()
	r.Use(ginpongo2.Pongo2("./"))

	r.GET("/", func(c *gin.Context) {
		c.Set("template", "index.html")
		c.Set("data", map[string]interface{}{"message": "Hello World!"})
	})

	r.GET("/none", func(c *gin.Context) {
		c.Set("template", "none.html")
	})

	r.GET("/invalidTemplate", func(c *gin.Context) {
		c.Set("template", 3)
	})

	r.GET("/invalidData", func(c *gin.Context) {
		c.Set("template", "index.html")
		c.Set("data", 3)
	})

	r.GET("/emptyData", func(c *gin.Context) {
		c.Set("template", "index.html")
		c.Set("data", nil)
	})

	e := r.Run(":8080")
	log.Println("stopping", e)
}
