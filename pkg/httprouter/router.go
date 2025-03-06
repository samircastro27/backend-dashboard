package httprouter

import (
	"context"
	"errors"
	"log"
	"strings"

	"net/http"

	"github.com/dapr/go-sdk/service/common"
)

type Router interface {
	HandleGet(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error))
	HandlePost(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error))
	HandlePut(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error))
	HandlePatch(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error))
	HandleDelete(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error))
	HandleOptions(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error))
	Handle(path string, method string, handler func(ctx context.Context, in Request) (out *common.Content, err error))
	StartServe()
}

type router struct {
	service common.Service
	routes  map[string]route
}

func NewRouter(service common.Service) *router {
	return &router{
		service: service,
		routes:  map[string]route{},
	}
}

func (r *router) HandleGet(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error)) {
	r.Handle(path, http.MethodGet, handler)
}
func (r *router) HandlePost(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error)) {
	r.Handle(path, http.MethodPost, handler)
}
func (r *router) HandlePut(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error)) {
	r.Handle(path, http.MethodPut, handler)
}
func (r *router) HandlePatch(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error)) {
	r.Handle(path, http.MethodPatch, handler)
}
func (r *router) HandleDelete(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error)) {
	r.Handle(path, http.MethodDelete, handler)
}
func (r *router) HandleOptions(path string, handler func(ctx context.Context, in Request) (out *common.Content, err error)) {
	r.Handle(path, http.MethodOptions, handler)
}

func (r *router) Handle(path string, method string, handler func(ctx context.Context, in Request) (out *common.Content, err error)) {
	route := route{
		Path: strings.ToLower(path),
	}
	if _, ok := r.routes[path]; ok {
		route = r.routes[path]
	}

	switch strings.ToUpper(method) {
	case http.MethodGet:
		route.Get = handler
	case http.MethodPost:
		route.Post = handler
	case http.MethodPut:
		route.Put = handler
	case http.MethodPatch:
		route.Patch = handler
	case http.MethodDelete:
		route.Delete = handler
	case http.MethodOptions:
		route.Options = handler
	}

	r.routes[path] = route

}

func (r *router) StartServe() {
	for _, route := range r.routes {
		r.service.AddServiceInvocationHandler(route.Path, func(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {

			switch in.Verb {
			case http.MethodGet:
				if route.Get != nil {
					return route.Get(ctx, &request{data: in.Data, queryString: in.QueryString})
				}
			case http.MethodPost:
				if route.Post != nil {
					return route.Post(ctx, &request{data: in.Data, queryString: in.QueryString})
				}
			case http.MethodPut:
				if route.Put != nil {
					return route.Put(ctx, &request{data: in.Data, queryString: in.QueryString})
				}
			case http.MethodPatch:
				if route.Patch != nil {
					return route.Patch(ctx, &request{data: in.Data, queryString: in.QueryString})
				}
			case http.MethodDelete:
				if route.Delete != nil {
					return route.Delete(ctx, &request{data: in.Data, queryString: in.QueryString})
				}
			case http.MethodOptions:
				if route.Options != nil {
					return route.Options(ctx, &request{data: in.Data, queryString: in.QueryString})
				}
			}

			err = errors.New("method not allowed")
			return out, err
		})
	}

	if err := r.service.Start(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error starting service: %v", err)
	}
}
