package main

import (
	"net/http"

	"github.com/Boofny/goLive"
	"github.com/Boofny/goLive/middleware"
)

func main() {
	e := goLive.Launch()

	e.Chain(
		middleware.CORS(),
		middleware.Logger(),
	)

	// Example get req
	e.GET("/ping", func(c *goLive.Context) error {
		return c.SendJSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	//example of a get request with path values
	e.GET("/user", func(c *goLive.Context) error {
		id := c.QueryGet("id")
		return c.SendSTRING(http.StatusOK, id)
	})
	
	//example of reading json from post request
	e.POST("/posting", func(c *goLive.Context) error {

		type User struct{ 
			Name string `json:"name"`
			Email string `json:"email"`
		}

		var data User
		err := c.ReadJSON(&data)
		if err != nil {
			return c.Error(http.StatusNotFound, "Error in /posting")
		}

		return c.PrettyJSON(http.StatusOK, map[string]any{
			"Name": "Hello " + data.Name,
			"Email": data.Email,
		})

	})

	v1 := e.GroupRoutes("/v1")

	v1.GET("/ping", func(c *goLive.Context) error {
		return c.SendSTRING(http.StatusOK, "pong v1")
	})

	e.StartServer(":8080")
}
