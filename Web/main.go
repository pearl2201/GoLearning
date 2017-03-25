package main

import (
	"Learn/web/db"

	"log"
	"net/http"
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
