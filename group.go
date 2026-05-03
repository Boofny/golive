package goLive

import "github.com/Boofny/goLive/middleware"

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

func (gr *RouteGroup) GET(path string, handle FunctionHandler) { //get request wrapper for simple usage
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return
	}
	gr.router.routes = append(gr.router.routes, route{
		method: "GET",
		pattern: gr.prefix + path,
		handler: handle,
	})
}

func (gr *RouteGroup)POST(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){ //put request wrapper
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return
	}
	gr.router.routes = append(gr.router.routes, route{
		method: "POST",
		pattern: gr.prefix + path,
		handler: handle,
	})
}

func (gr *RouteGroup)DELETE(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){ //DELETE request wrapper
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return// may need to add this to the others
	}
	gr.router.routes = append(gr.router.routes, route{
		method: "DELETE",
		pattern: gr.prefix + path,
		handler: handle,
	})
}

func (gr *RouteGroup)PUT(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){ //PUT request wrapper
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return// may need to add this to the others
	}
	gr.router.routes = append(gr.router.routes, route{
		method: "PUT",
		pattern: gr.prefix + path,
		handler: handle,
	})
}

func (gr *RouteGroup)Chain(mw ...middleware.Middleware){
	// g.middlewares = append(g.middlewares, middleware.Logger)
	gr.middlewares = append(gr.middlewares, mw...)
}
