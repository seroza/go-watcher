package main

import (
	"fmt"
	"log"
	"net/http"
)

// var templates = template.Must(template.ParseFiles(
// 	"templates/index.html",
// ))

// var validPath = regexp.MustCompile("^/(index|command)/([a-zA-Z0-9]+)$")
var ex Executor

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	var pagePath = "templates/" + tmpl + ".html"
	indexPageData, err := Asset(pagePath)

	if err != nil {
		log.Printf("Can't load %s", pagePath)
	}

	fmt.Fprintf(w, string(indexPageData))
}

func indexHandler(w http.ResponseWriter, r *http.Request, title string) {
	templateName := "index"
	renderTemplate(w, templateName)
}

func commandHandler(w http.ResponseWriter, r *http.Request, title string) {
	http.Redirect(w, r, "/", http.StatusFound)

	cmd, err := NewCommand(r)

	if err != nil {
		log.Println("Bad command received!")
		return
	}
	processCommand(cmd)
}

func processCommand(cmd Command) {
	ex.PushCommand(&cmd)
	log.Println(cmd)
}

func printRequestInfo(r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.Header)

	if err := r.ParseForm(); err != nil {
		log.Print("ParseForm error!")
	}

	for key, value := range r.PostForm {
		log.Printf("%s:%s", key, value)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, "")
	}
}

func main() {
	ex.Start()
	http.HandleFunc("/", makeHandler(indexHandler))
	http.HandleFunc("/command", makeHandler(commandHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
