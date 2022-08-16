package hw09structvalidator

import "strconv"

type intMinValidator struct{}

func (v intMinValidator) check(value int, condition string) bool {
	conditionInt, _ := strconv.Atoi(condition)
	return value >= conditionInt
}
