package ginpongo2

import (
	"github.com/flosch/pongo2"
	. "github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Pongo2() HandlerFunc {
	return func(c *Context) {
		c.Next()

		name, err := stringFromContext(c, "template")
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}

		data, _ := c.Get("data")

		template := pongo2.Must(pongo2.FromFile(name))
		err = template.ExecuteWriter(convertContext(data), c.Writer)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func stringFromContext(c *Context, input string) (string, error) {
	raw, ok := c.Get(input)
	if ok {
		strVal, ok := raw.(string)
		if ok {
			return strVal, nil
		}
	}
	return "", fmt.Errorf("No data for context variable: %s", input)
}

func convertContext(thing interface{}) pongo2.Context {
	if thing != nil {
		context, isMap := thing.(map[string]interface{})
		if isMap {
			return context
		}
	}
	return nil
}

func getContext(templateData interface{}, err error) pongo2.Context {
	if templateData == nil || err != nil {
		return nil
	}
	contextData, isMap := templateData.(map[string]interface{})
	if isMap {
		return contextData
	}
	return nil
}
