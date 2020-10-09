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
	//TODO : validate user

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
func StoreUrls(userid string, urls map[string]string) (bool, error) {
	res := true
	var Err error
	if len(urls) > 0 {
		// save the urls
		for name, url := range urls {
			url = strings.Trim(url, " ")
			// add http:// if url doesn't have it
			httpPrefix := strings.HasPrefix(url, "http://")
			httpsPrefix := strings.HasPrefix(url, "https://")
			if !(httpPrefix || httpsPrefix) {
				url = "http://" + url
			}
			ins, err := DB.Query("insert into urls (urlname, url, userid) values (?, ?, ?)", name, url, userid)
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
func FetchUrls(userid string) (urls map[string]string, err error) {
	var result *sql.Rows
	urls = make(map[string]string)
	result, err = DB.Query("Select urlname, url from urls where userid = ? group by urlname, url", userid)
	if err != nil {
		urls = nil
		return
	}

	for result.Next() {
		var url string
		var urlname string
		if err = result.Scan(&urlname, &url); err != nil {
			urls = nil
			break
		}
		if url != "" {
			if urlname == "" {
				urlname = url[:20]
			}
			urls[urlname] = url
		}
	}
	return
}
