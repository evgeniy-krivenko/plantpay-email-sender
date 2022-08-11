package qpye

import "log"

type Routes map[string]Handler

type Route struct {
	Pattern string
	Cb      Handler
}

type Router struct {
	Routes
}

func NewRouter(routes *[]Route) *Router {
	r := make(Routes, len(*routes))
	for _, v := range *routes {
		if _, ok := r[v.Pattern]; ok {
			log.Fatalf("route with pattern %s already exists", v.Pattern)
		}
		r[v.Pattern] = v.Cb
	}
	return &Router{r}
}
