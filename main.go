package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	password := "SQLpassword2023"
	dbName := "testdb"
	//properties := []string{"email", "first_name", "last_name"}
	//insertStudent(password, dbName, "student4@gmail.com", "Name4", "Last name4")
	//selectStudents(password, dbName)
	//updateStudent(password, dbName, properties[2], "Last name1", 1)
	//deleteStudent(password, dbName, 1)
	deleteAllStudents(password, dbName)
}

func insertStudent(password, dbName, email, first_name, last_name string) {

	db, err := sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/"+dbName)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO students(email, first_name, last_name) VALUES ('" + email + "', '" + first_name + "','" + last_name + "')"

	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}

type Student struct {
	Id         int
	Email      string
	First_Name string
	Last_Name  string
}

func selectStudents(password, dbName string) {
	db, err := sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/"+dbName)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM students")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var student Student
		err := res.Scan(&student.Id, &student.Email, &student.First_Name, &student.Last_Name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", student)
	}
}

func updateStudent(password, dbName, property, updateValue string, id int) {
	db, err := sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/"+dbName)
	ErrorCheck(err)

	defer db.Close()

	PingDB(db)

	stmt, e := db.Prepare("update students set " + property + "=? where id=?")
	ErrorCheck(e)

	// execute
	res, e := stmt.Exec(updateValue, id)
	ErrorCheck(e)

	a, e := res.RowsAffected()
	ErrorCheck(e)

	fmt.Println(a)
}

func deleteStudent(password, dbName string, id int) {
	db, e := sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/"+dbName)
	ErrorCheck(e)

	// close database after all work is done
	defer db.Close()

	PingDB(db)

	// delete data
	stmt, e := db.Prepare("delete from students where id=?")
	ErrorCheck(e)

	// delete 1st student
	res, e := stmt.Exec(id)
	ErrorCheck(e)

	// affected rows
	a, e := res.RowsAffected()
	ErrorCheck(e)

	fmt.Println(a) // 1
}

func deleteAllStudents(password, dbName string) {
	db, e := sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/"+dbName)
	ErrorCheck(e)

	// close database after all work is done
	defer db.Close()

	PingDB(db)

	// delete data
	stmt, e := db.Prepare("delete from students")
	ErrorCheck(e)

	// delete 1st student
	res, e := stmt.Exec()
	ErrorCheck(e)

	// affected rows
	a, e := res.RowsAffected()
	ErrorCheck(e)

	fmt.Println(a) // 1
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}
