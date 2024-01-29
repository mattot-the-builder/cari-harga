package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mattot-the-builder/go-csv/internal/app/models"
)

var layoutPath = filepath.Join("internal", "app", "views", "layout.html")
var premisePagePath = filepath.Join("internal", "app", "views", "premise.html")

func renderPremisePage(w http.ResponseWriter, data []models.Premise) {
	tmpl, err := template.New("").ParseFiles(premisePagePath, layoutPath)
	if err != nil {
		log.Println("Error parsing templates", err)
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

func GetAllPremise(w http.ResponseWriter, r *http.Request) {
	premiseList := models.GetPremiseList()
	renderPremisePage(w, premiseList)
}

func SearchPremise(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Wrong method request")
	}
	keyword := r.FormValue("search")
	log.Println(keyword)
	
	searchResult := models.SearchPremise(keyword)

	renderPremisePage(w, searchResult)
}
