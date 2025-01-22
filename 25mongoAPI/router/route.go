package route

import (
	"fmt"
	"mongoAPI/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	fmt.Println("hello")
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running on 4000 port"))
	})
	r.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.GetOneMovie).Methods("GET")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")

	return r
}
