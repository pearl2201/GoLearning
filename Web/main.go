package main

import (
	"Learn/web/db"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message = "complete task post"
	w.Write([]byte(message))
}

func ShowCompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message = "show complete task get"
	w.Write([]byte(message))
}

func ShowAllTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		context := db.GetTasks()
		message = context.Tasks[0].Title
		context.CSRF = "abcd"
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
		http.SetCookie(w, &cookie)
		templatesDir := "./templates/"
		var allFiles []string
		files, err := ioutil.ReadDir(templatesDir)
		if err != nil {
			log.Println(err)
			os.Exit(1) // No point in running app if templates aren't read
		}
		for _, file := range files {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				allFiles = append(allFiles, templatesDir+filename)
			}
		}

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		templates, err := template.ParseFiles(allFiles...)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		tmpl := templates.Lookup("home.html")

		tmpl.Execute(w, context)
	} else {
		message = "all pending task post"
	}
	w.Write([]byte(message))
}

func DeleteAllTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string = "delete task post"
	w.Write([]byte(message))
}

func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string = "show trash task get"
	w.Write([]byte(message))
}

//AddTaskFunc is used to handle the addition of new task, "/add" URL
func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	title := "random title"
	content := "random content"
	truth := db.AddTask(title, content)

	if truth != nil {
		log.Fatal("Error adding task")
	}
	w.Write([]byte("Adds task"))
}

func main() {

	PORT := "127.0.0.1:8080"
	log.Println("Running server on" + PORT)
	http.HandleFunc("/", ShowAllTaskFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	http.HandleFunc("/complete/", CompleteTaskFunc)
	http.HandleFunc("/completed/", ShowCompleteTaskFunc)
	http.HandleFunc("/delete/", DeleteAllTaskFunc)
	http.HandleFunc("/deleted/", ShowTrashTaskFunc)

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(PORT, nil))
}
