// Package goLive is the interface for accessing golives router methods
package goLive

// TODO: need to find out how to change all these package names from goLive to golive
import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Boofny/goLive/middleware"
)
	const banner =  `
 ██████╗  ██████╗ ██╗     ██╗██╗   ██╗███████╗██╗
██╔════╝ ██╔═══██╗██║     ██║██║   ██║██╔════╝██║
██║  ███╗██║   ██║██║     ██║██║   ██║█████╗  ██║
██║   ██║██║   ██║██║     ██║╚██╗ ██╔╝██╔══╝  ╚═╝
╚██████╔╝╚██████╔╝███████╗██║ ╚████╔╝ ███████╗██╗
 ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═══╝  ╚══════╝╚═╝
	`    

const (
	GET = http.MethodGet
	POST = http.MethodPost
	PUT = http.MethodPut
	DELETE = http.MethodDelete
	OPTIONS = http.MethodOptions
	PATCH = http.MethodPatch
	HEAD = http.MethodHead
)

var (
	ErrInvalidRedirectCode = errors.New("invalid redirect status code") 
	//will add more err code in future as this thing grows
)

type HandleFunc func(c *Context)error //custom handler defined for error handling

//GoLive dfined struct for starting server and chaining middleware
type GoLive struct{
	// Mux *http.ServeMux
	router *Router // using custom router for the reqs
	middlewares []middleware.Middleware
}

//Launch Method for starting the goLive session
func Launch()*GoLive{ 
	return &GoLive{
		router: &Router{
			routes: []route{},
		},
	}
}

func (g *GoLive) addRoute(method, path string, handle HandleFunc) {
	if path == "/favicon.ico" {
		return
	}
	parts := strings.Split(strings.Trim(path, "/"), "/")
	g.router.routes = append(g.router.routes, route{
		method:  method,
		pattern: path,
		parts:   parts,
		handler: handle,
	})
}

func (g *GoLive) GET(path string, handle HandleFunc) {
	g.addRoute(GET, path, handle)
}

func (g *GoLive) POST(path string, handle HandleFunc) {
	g.addRoute(POST, path, handle)
}

func (g *GoLive) DELETE(path string, handle HandleFunc) {
	g.addRoute(DELETE, path, handle)
}

func (g *GoLive) PUT(path string, handle HandleFunc) {
	g.addRoute(PUT, path, handle)
}

func (g *GoLive) OPTIONS(path string, handle HandleFunc) {
	g.addRoute(OPTIONS, path, handle)
}

//ServeStatic To serve a static file html txt png etc
func (g *GoLive) ServeStatic(urlPath, filepath string) error {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist %s", filepath)
	}

	g.router.routes = append(g.router.routes, route{
		method:  "GET",
		pattern: urlPath,
		handler: func(ctx *Context) error {
			http.ServeFile(ctx.Writer, ctx.Request, filepath)
			return nil
		},
	})
	return nil
}

//Chain use passes a variadic value of Middleware that is appended to the g.middlewares slice
func (g *GoLive)Chain(mw ...middleware.Middleware){
	// g.middlewares = append(g.middlewares, middleware.Logger)
	g.middlewares = append(g.middlewares, mw...)
}


//StartServer starts server with wrapped middleware and takes port ex: 8080
func (g *GoLive)StartServer(port string){

	stack := middleware.CreateStack(g.middlewares...)

	server := &http.Server{
		Addr:    port,
		Handler: stack(g.router), //where g.Mux is added after middleware chaining 
		// Handler: middleware.Logger(g.Mux), //where g.Mux is added after middleware chaining 
		// Handler: middleware.Logging(g.Mux), //this is where the output for Requests are
	}

	StartingDisaply(port)
	// for _, r := range g.router.routes { // could use for later testing 
	// 		fmt.Printf("[%s] %s\n", r.method, r.pattern)
	// }
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
	} else if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
	}
}

// StartingDisaply is now global so it can be used out side of golive
func StartingDisaply(port string){
	blue := "\033[34m"
	yellow := "\033[33m"
	reset := "\033[30m"
	fmt.Println(blue, banner)
	fmt.Print("\033[34m >>> \033[0m")
	fmt.Print("Server started successfully on port " +  yellow + port + reset)
	fmt.Println("\033[34m <<< \033[0m")
	fmt.Println("--------------------------------------------------")
}
