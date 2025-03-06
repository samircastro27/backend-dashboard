package httprouter

import (
	"context"

	"github.com/dapr/go-sdk/service/common"
)

type RouteHandle func(ctx context.Context, in Request) (out *common.Content, err error)

type route struct {
	Path    string
	Get     RouteHandle
	Post    RouteHandle
	Delete  RouteHandle
	Patch   RouteHandle
	Put     RouteHandle
	Options RouteHandle
}
