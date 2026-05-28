package goLive

import (
	"strings"

	"github.com/Boofny/goLive/middleware"
)

type RouteGroup struct {
	prefix string
	router *Router
	middlewares []middleware.Middleware
}

func (g *GoLive) GroupRoutes(prefix string) *RouteGroup{
	return &RouteGroup{
		prefix: prefix,
		router: g.router,
	}
}

func (gr *RouteGroup) addRoute(method, path string, handle HandleFunc) {
	if path == "/favicon.ico" {
		return
	}
	fullPath := gr.prefix + path
	parts := strings.Split(strings.Trim(fullPath, "/"), "/")
	gr.router.routes = append(gr.router.routes, route{
		method:  method,
		pattern: fullPath,
		parts:   parts,
		handler: handle,
	})
}

func (gr *RouteGroup) GET(path string, handle HandleFunc) {
	gr.addRoute(GET, path, handle)
}

func (gr *RouteGroup) POST(path string, handle HandleFunc) {
	gr.addRoute(POST, path, handle)
}

func (gr *RouteGroup) DELETE(path string, handle HandleFunc) {
	gr.addRoute(DELETE, path, handle)
}

func (gr *RouteGroup) PUT(path string, handle HandleFunc) {
	gr.addRoute(PUT, path, handle)
}

func (gr *RouteGroup) OPTIONS(path string, handle HandleFunc) {
	gr.addRoute(OPTIONS, path, handle)
}

func (gr *RouteGroup)Chain(mw ...middleware.Middleware){
	// g.middlewares = append(g.middlewares, middleware.Logger)
	gr.middlewares = append(gr.middlewares, mw...)
}
