package hw09structvalidator

import "strconv"

type intMaxValidator struct{}

func (v intMaxValidator) check(value int, condition string) bool {
	conditionInt, _ := strconv.Atoi(condition)
	return value <= conditionInt
}
