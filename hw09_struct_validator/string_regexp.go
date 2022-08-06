package hw09structvalidator

import (
	"regexp"
)

type stringRegexpValidator struct{}

func (v stringRegexpValidator) check(value interface{}, condition interface{}) bool {
	r := regexp.MustCompile(condition.(string))
	result := r.FindAllString(value.(string), -1)

	return result != nil && result[0] == value.(string)
}
