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
	r.HandleFunc("/premise", controllers.GetAllPremise).Methods("GET")
	r.HandleFunc("/premise", controllers.SearchPremise).Methods("POST")
	r.HandleFunc("/price", controllers.GetAllPrices).Methods("GET")
}
