package utils

import "strings"

// ValidateUserID : validate the userid
func ValidateUserID(userid string) (valid bool) {
	valid = true
	if len(userid) > 5 && strings.ContainsAny(userid, "~`!@#$%^&*()+-={}[]|\\:;\"'<,>?/") {
		valid = false
	}
	return
}
