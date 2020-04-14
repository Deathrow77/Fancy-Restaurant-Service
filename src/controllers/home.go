package controllers

import (
	"html/template"
	"net/http"
	"viewmodel"
)

// Storing the pointer to homeTempate in a struct
type home struct {
	homeTemplate *template.Template
}

// Route requests to home to handleHome function
func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
}

// Write the Base ViewModel to the Home Template
func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewBase()
	h.homeTemplate.Execute(w, vm)
}
