package ginpongo2

import (
	"github.com/flosch/pongo2"
	. "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Pongo2(templatePath string) HandlerFunc {
	return func(c *Context) {
		c.Next()

		templateName, templateNameError := c.Get("template")
		templateNameValue, isString := templateName.(string)

		if templateNameError == nil && isString {
			templateData, templateDataError := c.Get("data")
			var context = getContext(templateData, templateDataError)
			var template = pongo2.Must(pongo2.FromFile(templateNameValue))
			err := template.ExecuteRW(c.Writer, &context)
			if err != nil {
				http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			}
		}
	}
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
