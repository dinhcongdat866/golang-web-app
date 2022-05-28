package utils

import (
	"database/sql"
	"log"
)

type User struct {
	username string
}

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/user?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db

}

func SearchUserPwd(username string, password string) bool {
	db := ConnectDatabase()
	defer db.Close()

	results, err := db.Query("SELECT username FROM users WHERE username = ? AND pwd = ?", username, password)
	defer results.Close()

	if err != nil {
		log.Fatal(err)
		return false
	}

	for results.Next() {
		return true
	}

	return false
}

func SearchUser(username string) bool {
	db := ConnectDatabase()
	defer db.Close()

	results, err := db.Query("SELECT username FROM users WHERE username = ?", username)
	defer results.Close()

	if err != nil {
		log.Fatal(err)
		return false
	}

	for results.Next() {
		var user User

		err := results.Scan(&user.username)

		if err != nil {
			return false
		}
		return true
	}

	return false
}
