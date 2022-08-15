package hw09structvalidator

import (
	"strconv"
	"strings"
)

type intInValidator struct{}

func (v intInValidator) check(value int, condition string) bool {
	conditions := strings.Split(condition, ",")
	return containsInt(conditions, value)
}

func containsInt(s []string, e int) bool {
	for _, a := range s {
		val, _ := strconv.Atoi(a)
		if val == e {
			return true
		}
	}
	return false
}
