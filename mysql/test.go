package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/college?parseTime=true")
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err.Error())
	}
	fmt.Println("Connected")

	//create a new table
	{
		query := `

			CREATE TABLE student(
				id INT AUTO_INCREMENT,
        		name TEXT NOT NULL,
        		course TEXT NOT NULL,
        		created_at DATETIME,
        		PRIMARY KEY (id)
			);`
		if _, err := db.Exec(query); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Created table")
	}

	{
		//insert new entry

		name := "RAJU"
		course := "CSE"
		createdAt := time.Now()

		entry, err := db.Exec(`INSERT INTO student(name,course,created_at) VALUES(?,?,?)`, name, course, createdAt)

		if err != nil {
			fmt.Println(err.Error())
		}
		id, err := entry.LastInsertId()
		fmt.Println(id)
	}

	{
		//query single user

		var (
			id         int
			name       string
			course     string
			created_at time.Time
		)
		query := "SELECT id,name,course,created_at FROM student WHERE id = ?"

		if err := db.QueryRow(query, 1).Scan(&id, &name, &course, &created_at); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(id, name, course, created_at)
	}

	{ //to query all records
		type student struct {
			id         int
			name       string
			course     string
			created_at time.Time
		}

		query, err := db.Query("SELECT id,name,course,created_at FROM student")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer query.Close()

		var students []student

		for query.Next() {
			var u student
			err := query.Scan(&u.id, &u.name, &u.course, &u.created_at)
			if err != nil {
				fmt.Println(err.Error())
			}
			students = append(students, u)
		}

		fmt.Println("users:", students)

	}

}
