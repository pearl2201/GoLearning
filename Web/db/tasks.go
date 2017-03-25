package db

import (
	"Learn/web/types"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

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
	database.db, err = sql.Open("mysql", "pearl:hoanghien@/Pearl?parseTime=true")
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

//GetTasks get all tasks from database
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
		a := types.Task{Id: TaskID, Title: TaskTitle, Content: TaskContent, Created: TaskCreated.Format(time.UnixDate)[0:20]}
		task = append(task, a)
	}
	context = types.Context{Tasks: task}
	return context

}

//AddTask add task to database
func AddTask(title, content string) error {
	var err error
	query := "insert into task(title, content,created_date, last_modified_at,cat_id,task_status_id,user_id) values(?,?,now(), now(),?,?,?)"
	restoreSQL := database.prepare(query)
	tx := database.begin()
	_, err = tx.Stmt(restoreSQL).Exec(title, content, 1, 1, 1)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	} else {
		log.Println("insert successful")
		tx.Commit()
	}
	return err
}
