package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattot-the-builder/go-csv/internal/app/controllers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "route works")
	})
	r.HandleFunc("/item", controllers.HandleGetAllItem).Methods("GET")
	r.HandleFunc("/item", controllers.HandleSearchItem).Methods("POST")

	r.HandleFunc("/premise", controllers.HandleGetAllPremise).Methods("GET")
	r.HandleFunc("/premise", controllers.HandleSearchPremise).Methods("POST")

}
