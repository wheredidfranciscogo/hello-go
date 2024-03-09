package main

import (
	// "fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	// })
	// serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start server
	http.ListenAndServe(":80", nil)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// func main() {
// 	// Extension 1: Serve static files
// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	// Extension 2: Implement routing
// 	http.HandleFunc("/", handleRoot)
// 	http.HandleFunc("/about", handleAbout)

// 	// Extension 3: Render dynamic HTML pages using Go templates
// 	http.HandleFunc("/dynamic", handleDynamic)

// 	// Start the server
// 	fmt.Println("Server listening on port 8080...")
// 	http.ListenAndServe(":8080", nil)
// }

// func handleRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
// }

// func handleAbout(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is the about page.")
// }

// func handleDynamic(w http.ResponseWriter, r *http.Request) {
// 	// Extension 3: Render dynamic HTML page using Go templates
// 	tmpl, err := template.New("index").Parse(`<html><body><h1>Hello, {{.Name}}!</h1></body></html>`)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	data := struct{ Name string }{"John"}
// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
