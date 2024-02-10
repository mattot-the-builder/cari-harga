package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mattot-the-builder/go-csv/internal/app/models"
)

var itemPagePath = filepath.Join("internal", "app", "views", "item.html")

func renderItemPage(w http.ResponseWriter, data []models.Item) {
	tmpl, err := template.New("").ParseFiles(itemPagePath, layoutPath)
	if err != nil {
		log.Println("Error parsing templates", err)
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

func HandleGetAllItem(w http.ResponseWriter, r *http.Request) {
	itemList := models.GetItemList()
	renderItemPage(w, itemList)
}

func HandleSearchItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Wrong method request")
	}
	keyword := r.FormValue("search")
	log.Println(keyword)
	
	searchResult := models.SearchItem(keyword)

	renderItemPage(w, searchResult)
}
