package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattot-the-builder/go-csv/internal/app/routes"
)

func main() {
	// initialTime := time.Now()

	// data := utils.ReadCsvFile("data/lookup_premise.csv")


	// controllers.PrintPremiseList(premiseList)

	// fmt.Printf("Total rows iterated: %v\n", len(premiseList))
	// fmt.Printf("Time taken: %v\n", time.Now().Sub(initialTime))

	// premiseList := models.GetPremiseList()
	// fmt.Println(premiseList)

	r := mux.NewRouter()
	// register routes into mux router
	routes.RegisterRoutes(r)

	// start server
	http.Handle("/", r)
	log.Println("Starting server at localhost:8080")
	r.Handle("/static/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", r)

}
