package main

import (
	"log"
	"net/http"
)





func aboutPage(w http.ResponseWriter, r *http.Request) {
	
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	
	http.ServeFile(w, r, "static/contact.html")
}

func main() {

	
	
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
