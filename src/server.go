package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"viewmodel"
)

//
// type myHandler struct {
// 	name string
// }
//
// func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(fmt.Sprintf("Hello %v !!", mh.name)))
// }

func main() {
	// Render all templates for one
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Store the required file
		requiredFile := r.URL.Path[1:]
		// Run a lookup for the required template
		t := templates[requiredFile+".html"] // Lookup is not done since templates variable in now a map of templates
		context := viewmodel.NewBase()
		if t != nil {
			t.Execute(w, context)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	// fmt.Println(os.Getwd())
	// Handle CSS and Image requirements for basic styling
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/fonts/", http.FileServer(http.Dir("public")))
	// Run the Server
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() map[string]*template.Template {
	// Make a map of templates
	result := make(map[string]*template.Template)
	// Create a base Path variable
	const basePath = "templates"
	// load the common templates into a layout variable
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	// Open the Content Directory
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Error in opening the content folder - " + err.Error())
	}
	// Read all the contents of the Directory
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Error in Reading the Content Folder - " + err.Error())
	}
	// Iterate over the contents of the folder
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		fmt.Println(fi.Name())
		if err != nil {
			panic("Error in Reading the template - " + err.Error())
		}
		// Read the content of the File
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Error in Reading the Contents of the template - " + err.Error())
		}
		// Close the file reader
		f.Close()
		// Clone the layout template into a new template variable
		tmp1 := template.Must(layout.Clone())
		// Add the content to the new template variable
		_, err = tmp1.Parse(string(content))
		if err != nil {
			panic("Error Parsing the Content of the templates - " + err.Error())
		}
		// Add the template to the map wrt the corresponding file name
		result[fi.Name()] = tmp1
	}
	return result
}
