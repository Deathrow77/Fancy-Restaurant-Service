package controllers

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	menuController menu
)

func Startup(templates map[string]*template.Template) {
	// Call the Home Controller and map the respective home template
	homeController.homeTemplate = templates["index.html"]
	menuController.menuTemplate = templates["menu.html"]
	homeController.registerRoutes()
	menuController.registerRoutes()
	// Handle Static File Loading
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/fonts/", http.FileServer(http.Dir("public")))
}
