package db

import (
	"Learn/web/types"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"time"
)

var database Database
var taskStatus map[string]int
var err error

type Database struct {
	db *sql.DB
}

// Begin a transaction

func (db1 Database) begin() (tx *sql.Tx) {
	tx, err := db1.db.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	return tx

}

func (db1 Database) prepare(q string) (stmt *sql.Stmt) {
	stmt, err := db1.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil
	}
	return stmt
}

func (db Database) query(q string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.db.Query(q, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rows, nil
}

func init() {
	database.db, err = sql.Open("mysql", "pearl:hoanghien@/Pearl")
	taskStatus = map[string]int{"COMPLETE": 1, "PENDING": 2, "DELETED": 3}
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	database.db.Close()
}

/*func getTasks(username, status, category string) (types.Context, error) {
	var tasks []Task
	var navigation string
	var search string
	var message string
	var referer string
	var csrf string
	//	var category []types.CategoryCount

}
*/

func GetTasks() types.Context {
	var task []types.Task
	var context types.Context
	var TaskID int
	var TaskTitle string
	var TaskContent string
	var TaskCreated time.Time
	var getTaskSql string
	getTaskSql = "select id, title, content, created_date from task;"
	rows, err := database.query(getTaskSql)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&TaskID, &TaskTitle, &TaskContent, &TaskCreated)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
		TaskCreated = TaskCreated.Local()
		a := types.Task{Id: TaskID, Title: TaskTitle, Content: TaskContent, Created: TaskCreated.String()}
		task = append(task, a)
	}
	context = types.Context{Tasks: task}
	return context

}
