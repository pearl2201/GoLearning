package types

import "html/template"

/*
Packages type is used to store the context struct whish is passed while template are executed
*/

//Task is the struct used to identify tasks
type Task struct {
	Id               int           `json:"id"`
	Title            string        `json:"title"`
	Content          string        `json:"content"`
	ContentHTML      template.HTML `json:"content_html"`
	Created          string        `json:"created"`
	Priority         string        `json:"priority"`
	Category         string        `json:"category"`
	Referer          string        `json:"referer,omitempty"`
	Comments         []string      `json:"comment,omitempty"`
	IsOverdue        bool          `json:"isoverdue,omitempty"`
	IsHidden         int           `json:"ishidden, omitempty"`
	CompletedMessage string        `json:"completedMessage,omitempty"`
}

type Tasks []Task

// Comment is the struct used to populate comments per task
type Comment struct {
	ID       int
	Content  string
	Created  string
	Username string
}

// Context is the struct passed to template
type Context struct {
	Tasks      []Task
	Navigation string
	Search     string
	Message    string
	CSRF       string
	Category   []CategoryCount
	Referre    string
}

//CategoryCount is the struct used to populate the sidebar
type CategoryCount struct {
	Name  string
	Count int
}

// Status is the json struct used to return
type Status struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

//Category is the structure of the category table
type Category struct {
	ID      int    `json:"category_id"`
	Name    string `json:"category_name"`
	Created string `json:"category_created"`
}

//Categories will show
type Categories []Category
