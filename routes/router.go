package routes

import (
	"fmt"
	"net/http"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	http.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server is listining")
	})

	return mux
}
