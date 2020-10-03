package models

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// DB instance
var DB *sql.DB

// InitDB : initialize the db and connect
func InitDB() {
	database, err := sql.Open("mysql", "jiraiya:Shivi<323@tcp(database-1.caqh2nel7qhl.us-east-2.rds.amazonaws.com:3306)/cumul")

	if err != nil {
		panic(err.Error())
	}

	DB = database
}

// AddUser Store the user
func AddUser(userid string) (added bool, err error) {
	added = true
	stmt, err := DB.Prepare("INSERT INTO user (userid) values (?)")
	if err != nil {
		added = false
	}
	defer stmt.Close()
	_, err = stmt.Exec(userid)
	if err != nil {
		added = false
	}
	return
}

// UserExists check whether user is present in the DB
func UserExists(userid string) (exists bool) {
	var user string
	exists = true
	err := DB.QueryRow("select userid from user where userid=?", userid).Scan(&user)
	if err != nil || user == "" {
		exists = false
		return
	}
	return
}

// StoreUrls store all the urls for the given user
func StoreUrls(userid string, urls []string) (bool, error) {
	res := true
	var Err error
	if len(urls) > 0 {
		// save the urls
		for _, u := range urls {
			u = strings.Trim(u, " ")
			ins, err := DB.Query("insert into urls (url, userid) values (?, ?)", u, userid)
			if err != nil {
				res = false
				Err = err
				break
			}
			ins.Close()
		}
	}
	return res, Err
}

// FetchUrls : fetches all the urls for the given userid
func FetchUrls(userid string) (urls []string, err error) {
	var result *sql.Rows
	result, err = DB.Query("Select distinct(url) from urls where userid = ?", userid)
	if err != nil {
		urls = nil
		return
	}

	for result.Next() {
		var url string
		if err = result.Scan(&url); err != nil {
			urls = nil
			break
		}

		urls = append(urls, url)
	}
	return
}
