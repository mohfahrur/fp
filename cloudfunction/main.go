package helloworld

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gin-gonic/gin"
)

func init() {
	// functions.Typed("HelloHTTP", helloHTTP)
	router := gin.Default()
	router.POST("/hello", hello)
	functions.HTTP("HelloGin", router.ServeHTTP)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func helloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	fmt.Fprintf(w, "Hi, %s!", html.EscapeString(d.Name))
}
