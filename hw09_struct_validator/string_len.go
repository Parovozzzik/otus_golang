package hw09structvalidator

import "strconv"

type stringLenValidator struct{}

func (v stringLenValidator) check(value interface{}, condition interface{}) bool {
	length, _ := strconv.Atoi(condition.(string))
	return len(value.(string)) == length
}
