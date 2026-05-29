## GoLive!
### Fast web frame work for golang built on std lib http

## Installation 
### Get started by downloading this module
#### Need Golang 1.25.1+
```bash
go get github.com/Boofny/golive@latest
```

### Simple Example
```go
package main

import (
	"net/http"
	"github.com/Boofny/golive"
	"github.com/Boofny/golive/middleware"
)

func main() {
	e := golive.Launch()

	e.Chain(middleware.CORS())

	e.GET("/ping", func(c *golive.Context)error{
		return c.SendSTRING(http.StatusOK, "pong") //send out your data
	})

	e.StartServer(":8080")
}
```

## Try it out!
### In the terminal try.
```bash
curl -X GET localhost:8080/hello
```
### Or in your browser 
```bash
http://localhost:8080/hello
```
### Check out the docker image 
[Docker hub](https://hub.docker.com/repository/docker/boofny/golive-docker/general) 

### Or docker one liner for a quick demo
```bash
docker run -p 8000:8000 -e PORT=8000 boofny/golive-docker
```
