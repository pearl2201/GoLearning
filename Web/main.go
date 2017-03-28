package main

import (
	"Learn/web/db"
	"Learn/web/views"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
		views.HomeTemplate.Execute(w, context)
	} else {
		message = "all pending task post"
		w.Write([]byte(message))
	}

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
	if r.Method == "GET" {
		r.ParseForm()
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println(err)
		}
		taskPriority, priorityErr := strconv.Atoi(r.FormValue("priority"))
		if priorityErr != nil {
			log.Println("Unable to convert priority to integer")

		}
		priorityList := []int{1, 2, 3}
		for _, priority := range priorityList {
			if taskPriority != priority {
				log.Println("incorrect priority sent")
				taskPriority = 1
			}
		}
		title := template.HTMLEscapeString(r.Form.Get("title"))
		content := template.HTMLEscapeString(r.Form.Get("content"))
		formToken := template.HTMLEscapeString(r.Form.Get("CSRFToken"))
		truth := db.AddTask(title, content)
		cookie, _ := r.Cookie("csrftoken")
		if formToken == cookie.Value {
			if handler != nil {
				r.ParseMultipartForm(32 << 20)
				defer file.Close()
				f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					log.Println(err)
					return
				}
				defer f.Close()
				io.Copy(f, file)
				filelink := "<br> <a href=/files/" + handler.Filename + ">" + handler.Filename + "</a>"
				content = content + filelink
			}
			truth := db.AddTask(title, content, taskPriority)
			if truth != nil {
				message = "Error adding task"
				log.Fatal("Error adding task to db")
			} else {
				message = "Task added"
				log.Println("added task to db")
			}
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			log.Fatal("csrf mismatch")
		}

	} else {
		message = "Method not allowed"
		http.Redirect(w, r, "/", http.StatusFound)
	}

	w.Write([]byte("Adds task"))
}

func main() {
	views.PopulateTemplates()
	PORT := "127.0.0.1:8080"
	log.Println("Running server on: " + PORT)
	http.HandleFunc("/", ShowAllTaskFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	http.HandleFunc("/complete/", CompleteTaskFunc)
	http.HandleFunc("/completed/", ShowCompleteTaskFunc)
	http.HandleFunc("/delete/", DeleteAllTaskFunc)
	http.HandleFunc("/deleted/", ShowTrashTaskFunc)

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(PORT, nil))
}
