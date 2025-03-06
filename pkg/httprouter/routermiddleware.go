package httprouter

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dapr/go-sdk/service/common"
)

type MiddlewareFunc func(next RouteHandle) RouteHandle

type Middleware struct {
	middlewares []MiddlewareFunc
}

func (m *Middleware) Use(middleware MiddlewareFunc) {
	m.middlewares = append(m.middlewares, middleware)
}

func (m *Middleware) Execute(handler RouteHandle) RouteHandle {
	h := handler
	for i := len(m.middlewares) - 1; i >= 0; i-- {
		h = m.middlewares[i](h)
	}
	return h
}

// Extender el router existente para incluir middleware
type RouterWithMiddleware struct {
	service    common.Service
	routes     map[string]route
	middleware []MiddlewareFunc
}

// NewRouterWithMiddleware creates a new router with middleware support
func NewRouterWithMiddleware(service common.Service) *RouterWithMiddleware {
	return &RouterWithMiddleware{
		service:    service,
		routes:     make(map[string]route),
		middleware: make([]MiddlewareFunc, 0),
	}
}

// Use adds middleware to the router
func (r *RouterWithMiddleware) Use(middleware MiddlewareFunc) {
	r.middleware = append(r.middleware, middleware)
}

func (r *RouterWithMiddleware) applyMiddleware(handler RouteHandle) RouteHandle {
	h := handler
	// Apply middleware in reverse order
	for i := len(r.middleware) - 1; i >= 0; i-- {
		h = r.middleware[i](h)
	}
	return h
}

func (r *RouterWithMiddleware) Handle(path string, method string, handler RouteHandle) {
	route := route{
		Path: strings.ToLower(path),
	}

	if existingRoute, ok := r.routes[path]; ok {
		route = existingRoute
	}

	// Apply middleware to the handler
	wrappedHandler := r.applyMiddleware(handler)

	switch strings.ToUpper(method) {
	case http.MethodGet:
		route.Get = wrappedHandler
	case http.MethodPost:
		route.Post = wrappedHandler
	case http.MethodPut:
		route.Put = wrappedHandler
	case http.MethodPatch:
		route.Patch = wrappedHandler
	case http.MethodDelete:
		route.Delete = wrappedHandler
	case http.MethodOptions:
		route.Options = wrappedHandler
	}

	r.routes[path] = route
}

func (r *RouterWithMiddleware) HandleGet(path string, handler RouteHandle) {
	r.Handle(path, http.MethodGet, handler)
}

func (r *RouterWithMiddleware) HandlePost(path string, handler RouteHandle) {
	r.Handle(path, http.MethodPost, handler)
}

func (r *RouterWithMiddleware) HandlePut(path string, handler RouteHandle) {
	r.Handle(path, http.MethodPut, handler)
}

func (r *RouterWithMiddleware) HandlePatch(path string, handler RouteHandle) {
	r.Handle(path, http.MethodPatch, handler)
}

func (r *RouterWithMiddleware) HandleDelete(path string, handler RouteHandle) {
	r.Handle(path, http.MethodDelete, handler)
}

func (r *RouterWithMiddleware) HandleOptions(path string, handler RouteHandle) {
	r.Handle(path, http.MethodOptions, handler)
}

type contextKey string

const (
	httpMethodKey contextKey = "http-method"
	pathKey       contextKey = "path"
)

func (r *RouterWithMiddleware) StartServe() {
	for _, route := range r.routes {
		r.service.AddServiceInvocationHandler(route.Path, func(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
			req := &request{
				data:        in.Data,
				queryString: in.QueryString,
				method:      in.Verb,
			}

			ctx = context.WithValue(ctx, httpMethodKey, in.Verb)
			ctx = context.WithValue(ctx, pathKey, route.Path)

			var handler RouteHandle
			switch in.Verb {
			case http.MethodGet:
				handler = route.Get
			case http.MethodPost:
				handler = route.Post
			case http.MethodPut:
				handler = route.Put
			case http.MethodPatch:
				handler = route.Patch
			case http.MethodDelete:
				handler = route.Delete
			case http.MethodOptions:
				handler = route.Options
			}

			if handler == nil {
				return nil, errors.New("method not allowed")
			}

			return handler(ctx, req)
		})
	}

	if err := r.service.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error starting service: %v", err)
	}
}
