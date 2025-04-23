package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rahimuj570/go_net-http_BLOG_REST_API/config"
	"github.com/rahimuj570/go_net-http_BLOG_REST_API/routes"
)

func main() {
	//Load Env
	config.LoadEnv()
	router := routes.SetupRouter()
	config.InitDB()

	// Configuring Server
	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	//Listining Server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server Error: ", err)
	}
}
