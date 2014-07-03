# ginpongo2

The `github.com/ngerakines/ginpongo2` package provides a middleware that can be used to render pongo2 templates.

## Example

To use this, first ensure that the middleware is being referenced by gin using `Use`. Then inide of your handler, use the context `Set` methods to set the "template" and "data" variables.

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngerakines/ginpongo2"
	"log"
)

func main() {
	r := gin.Default()
	r.Use(ginpongo2.Pongo2())

	r.GET("/", func(c *gin.Context) {
		c.Set("template", "index.html")
		c.Set("data", map[string]interface{}{"message": "Hello World!"})
	})

	e := r.Run(":8080")
}

```
