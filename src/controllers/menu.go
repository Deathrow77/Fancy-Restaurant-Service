package controllers

import (
	"html/template"
	"net/http"
	"viewmodel"
)

// Storing the menuTemplate pointer in a struct
type menu struct {
	menuTemplate *template.Template
}

// Pass all requests to /menu to handleMenu Handler Function
func (m menu) registerRoutes() {
	http.HandleFunc("/menu", m.handleMenu)
}

// map the menu viewmodel to the menu template
func (m menu) handleMenu(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.CreateMenu()
	m.menuTemplate.Execute(w, vm)
}
