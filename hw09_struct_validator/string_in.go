package hw09structvalidator

import "strings"

type stringInValidator struct{}

func (v stringInValidator) check(value interface{}, condition interface{}) bool {
	conditions := strings.Split(condition.(string), ",")
	return containsStr(conditions, value.(string))
}

func containsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
