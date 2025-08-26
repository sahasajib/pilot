package cmd

import (
	"net/http"

	"github.com/sahasajib/pilot/handler"
	"github.com/sahasajib/pilot/middleware"
)

func InitRoute(route *http.ServeMux, manager *middleware.Manager){
	route.Handle("GET /user", manager.With(http.HandlerFunc(handler.Users)))
	route.Handle("POST /user", manager.With(http.HandlerFunc(handler.CreateUser)))
}