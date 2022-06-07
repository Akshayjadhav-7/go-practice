package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/customers?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	{
		//create new table

		table := `
				CREATE TABLE customers(
				id INT AUTO_INCREMENT,
				name TEXT NOT NULL,
				address TEXT NOT NULL,
				phone INT NOT NULL,
				PRIMARY KEY (id)
			);`

		if _, err := db.Exec(table); err != nil {
			fmt.Println(err.Error())

		}
		fmt.Println("Table created")
	}

	// {
	// 	//insert 1st entry

	// 	name := "AJ"
	// 	address := "BMT"
	// 	phone := 12345444

	// 	entry, err := db.Exec(`INSERT INTO customers(name,address,phone) VALUES(?,?,?)`, name, address, phone)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	id, err := entry.LastInsertId()
	// 	fmt.Println(id)
	// }

	{
		//query single entry

		var (
			id      int
			name    string
			address string
			phone   int
		)

		query := "SELECT id,name,address,phone FROM customers WHERE id = ?"

		if err := db.QueryRow(query, 1).Scan(&id, &name, &address, &phone); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(id, name, address, phone)
	}

	{
		//query all records

		type user struct {
			id      int
			name    string
			address string
			phone   int
		}

		query, err := db.Query("SELECT id,name,address,phone FROM customers")
		if err != nil {
			fmt.Println(err.Error())
		}

		defer query.Close()

		var users []user

		for query.Next() {
			var u user

			err := query.Scan(&u.id, &u.name, &u.address, &u.phone)
			if err != nil {
				fmt.Println(err.Error())
			}

			users = append(users, u)
		}
		fmt.Println(users)
	}

	// {

	// 	// update the record
	// 	id := 1
	// 	name := "Sanjay"
	// 	address := "Mumbai"
	// 	phone := 654321

	// 	update, err := db.Exec("update customers set name=?,address=?,phone=? where id=?", name, address, phone, id)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	value, _ := update.RowsAffected()
	// 	fmt.Println(value)

	// }

	{
		//delete the record
		id := 1
		delete, err := db.Exec("delete from customers where id=?", id)
		if err != nil {
			fmt.Println(err.Error())
		}
		value, _ := delete.RowsAffected()
		fmt.Println(value)
	}
}

//update and delete
