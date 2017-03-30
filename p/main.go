package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	message := "check"
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
	w.Write([]byte(message))
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		firstname := template.HTMLEscapeString(r.Form.Get("firstname"))
		lastname := template.HTMLEscapeString(r.Form.Get("lastname"))
		log.Println("ahihi: " + firstname + " - " + lastname)
		w.Write([]byte("Add success"))

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/log/", add)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
