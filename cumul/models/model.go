package models

import "time"

type User struct {
	Userid  string    `json:"userid"`
	Paid    bool      `json:"paid"`
	Created time.Time `json:"created"`
}

type Url struct {
	Urlid   int       `json:"urlid"`
	Userid  string    `json:"userid"`
	URL     string    `json:"url"`
	Created time.Time `json:"created"`
}
