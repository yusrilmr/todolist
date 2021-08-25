package main

import (
	"github.com/urfave/negroni"
	"github.com/yusrilmr/todolist/backend/common"
	"github.com/yusrilmr/todolist/backend/routers"
	"log"
	"net/http"
)

func main() {
	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}