package main

import (
	"github.com/rahimuj570/go_net-http_BLOG_REST_API/config"
)

func main() {
	//Load Env
	config.LoadEnv()
	// router := routes.SetupRouter()
	// config.InitDB()

	// // Configuring Server
	// server := &http.Server{
	// 	Addr:         ":" + os.Getenv("PORT"),
	// 	Handler:      router,
	// 	ReadTimeout:  time.Second * 10,
	// 	WriteTimeout: time.Second * 10,
	// }

	// //Listining Server
	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Fatal("Server Error: ", err)
	// }

	// var inter Inter=Impl{}
	// var inter Impl=i
}

type Inter interface {
	M1() string
	M2() string
}

type Impl struct{}

func (i Impl) M1() string {
	return ""
}
