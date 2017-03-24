package db

import (
	"Learn/Web/types"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var database Database
var taskStatus map[string]int
var err Error

type Database struct {
	db *sql.DB
}

// Begin a transaction

func (db Database) begin() (tx *sql.Tx) {
	tx, err := db.db.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	return tx

}

func (db Database) prepare(q string) (stmt *sql.Stmt) {
	stmt, err := db.db.prepare(q)
	if err != nil {
		log.Println(err)
		return nil
	}
	return stmt
}

func (db Database) query(q string, args ...interface{}) (row *sql.Rows) {
	rows, err := db.db.Query(q, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}

func init() {
	database.db, err = sql.Open("mysql", "pearl:hoanghien@/Pearl")
	taskStatus = map[string]int{"COMPLETE": 1, "PENDING": 2, "DELETED": 3}
	if err != nil {
		log.Fatal(err)
	}
}

func close() {
	database.db.Close()
}

func getTasks(username, status, category string) (types.Context, error) {
	var tasks []Task
	var navigation string
	var search string
	var message string
	var referer string
	var csrf string
	var category []types.CategoryCount
}
