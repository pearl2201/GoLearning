package main

import (
	"Learn/web/db"
	"db"
	"log"
	"net/http"
)

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string = "complete task post"
	w.Write([]byte(message))
}

func ShowCompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string = "show complete task get"
	w.Write([]byte(message))
}

func ShowAllTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "all pending task get"
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

func AddTaskFunc(w http.ResponseWriter, r *http.Request) {

}

func AddTask(title, content, category string, taskPriority int, username string, hidden int) error {
	log.Println("Add task: start function")
	var err error

	return err
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
