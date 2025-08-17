package cmd

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/sahasajib/pilot/global_router"
	"github.com/sahasajib/pilot/middleware"
)

func Serve(){
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)

	route := http.NewServeMux()
	InitRoute(route, manager)
	globalHandler := global_router.GlobalRouter(route)
	slog.Info("Starting server on port 8080")

	err := http.ListenAndServe(":8080",globalHandler)
	if err != nil{
		fmt.Println("Error starting server:", err)
	}
}