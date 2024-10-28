package server

import (
	"github.com/congmanh18/XuyenXam/OrderService/handler"
	"github.com/congmanh18/lucas-core/transport/http/method"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

func Routes(handler handler.OrderHandler) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/order",
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: handler.HandleCreate,
				},
			},
		},
	}
}
