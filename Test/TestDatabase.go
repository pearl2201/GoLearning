package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	db, err := sql.Open("mysql", "pearl:hoanghien@/Pearl")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	} else {
		fmt.Println("connect success")
	}
	defer db.Close()

	/*_, err = db.Exec("CREATE DATABASE Pearl")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	} else {
		fmt.Println("create success")
	}*/
	/*fmt.Println("create table user")
	_, err = db.Exec("Create table user(id integer primary key auto_increment,username varchar(100),password varchar(1000),email varchar(100))")
	checkError(err)
	fmt.Println("insert into user value")
	_, err = db.Exec("Insert into user values(1,'pearl2201','hoanghien','nguyenanhngoc.ftu@gmail.com')")
	checkError(err)*/
	/*fmt.Println("create table category")
	_, err = db.Exec("Create table category (id integer primary key auto_increment, name var char(1000) not null, user_id integer, foreign key(user_id) references user(id) )")
	checkError(err)*/
	/*fmt.Println("create table status")
	_, err = db.Exec("create table status(id integer primary key auto_increment, status varchar(50) not null)")
	checkError(err)
	fmt.Println("create table task")
	_, err = db.Exec("create table task(id integer primary key auto_increment, title varchar(100), content text, created_date timestamp, last_modified_at timestamp, finish_date timestamp,priority integer, cat_id integer not null, task_status_id integer not null, due_date integer, user_id integer not null, hide integer, foreign key(cat_id) references category(id), foreign key(task_status_id) references status(id))")
	fmt.Println("create table comment")
	_, err = db.Exec("create table comment(id integer primary key auto_increment, content text, taskID int not null, created datetime ,user_id integer not null,foreign key(taskID) references task(id), foreign key(user_id) references user(id))")
	checkError(err)
	fmt.Println("create table files")
	_, err = db.Exec("Create table files(name varchar(1000) not null, autoName varchar(255) not null, created_date timestamp, user_id int not null, foreign key(user_id) references user(id))")
	checkError(err)*/
	/*fmt.Println("insert category")
	_, err = db.Exec("Insert into category values(1,'TaskApp',1)")
	checkError(err)
	fmt.Println("insert stt")
	_, err = db.Exec("Insert into status values(1,'COMPLETE')")
	checkError(err)
	_, err = db.Exec("Insert into status values(2,'PENDING')")
	checkError(err)*/
	/*_, err = db.Exec("Insert into status values(3,'FINISHED')")
	checkError(err)
	fmt.Println("insert task")
	_, err = db.Exec("Insert into task values(1,'Publish on github','Publish the source of tasks and picsort on github','2015-11-12 15:30:59','2015-11-21 14:19:22','2015-11-17 17:02:18',3,1,1,NULL,1,0)")
	checkError(err)*/
	_, err = db.Exec("Insert into task values(4,'gofmtall','The idea is to run gofmt -w file.go on every go file in the listing, *Edit turns out this is is difficult to do in golang **Edit barely 3 line bash script. ','2015-11-12 16:58:31','2015-11-14 10:42:14','2015-11-13 13:16:48',3,1,1,NULL,1,0)")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal Error: %s", err)
		os.Exit(2)
	}

}
