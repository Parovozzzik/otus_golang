package hw09structvalidator

import (
	"errors"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	result := ""
	for _, err := range v {
		result = result + err.Field + ": " + err.Err.Error()
	}
	return result
}

var (
	ErrIntIn  = errors.New("unsupported value")
	ErrIntMax = errors.New("too big value")
	ErrIntMin = errors.New("too small value")
	ErrStrIn  = errors.New("unsupported value")
	ErrStrLen = errors.New("invalid value size")
	ErrStrReg = errors.New("invalid value format")
)

var Setting = map[string][]string{
	"string": {"len", "regexp", "in"},
	"int":    {"min", "max", "in"},
}

func Validate(v interface{}) ValidationErrors {
	var errors ValidationErrors

	st := reflect.TypeOf(v)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if condition, ok := field.Tag.Lookup("validate"); ok {
			if condition == "" {
				continue
			}
			conditions := strings.Split(condition, "|")

			elemType := getElemType(&field)
			if !inSetting(elemType) {
				continue
			}

			values := getValues(v, field.Name)
			if values.Kind() == reflect.Slice {
				valueType := values.Type().Elem().String()
				validateSlice(values.Interface(), conditions, valueType, field.Name, &errors)
			} else {
				valueType := values.Type().String()
				validateValue(values.Interface(), conditions, valueType, field.Name, &errors)
			}
		}
	}

	return errors
}

func getValues(v interface{}, field string) reflect.Value {
	structValue := reflect.ValueOf(v)
	valuesElem := structValue
	if structValue.Kind() == reflect.Ptr {
		valuesElem = structValue.Elem()
	}
	return valuesElem.FieldByName(field)
}

func getElemType(field *reflect.StructField) string {
	if field.Type.Kind() == reflect.Slice {
		return field.Type.Elem().Kind().String()
	}
	return field.Type.Kind().String()
}

func inSetting(varType string) bool {
	_, ok := Setting[varType]
	return ok
}

func validateSlice(values interface{}, conditions []string, valueType string, field string, errors *ValidationErrors) {
	if valueType == "int" {
		for _, value := range values.([]int) {
			validateValue(value, conditions, valueType, field, errors)
		}
	}

	if valueType == "string" {
		for _, value := range values.([]string) {
			validateValue(value, conditions, valueType, field, errors)
		}
	}
}

func validateValue(value interface{}, conditions []string, valueType string, field string, errors *ValidationErrors) {
	res := true
	for _, subCondition := range conditions {
		conditionType := strings.Split(subCondition, ":")[0]
		conditionValue := strings.Split(subCondition, ":")[1]

		if valueType == "string" {
			res = validateStr(value.(string), conditionType, conditionValue)
		}

		if valueType == "int" {
			res = validateInt(value.(int), conditionType, conditionValue)
		}

		if !res {
			*errors = append(*errors, ValidationError{Field: field, Err: getTypeError(valueType, conditionType)})
			break
		}
	}
}

func validateInt(varA int, rule string, varB string) bool {
	switch rule {
	case "in":
		validator := intInValidator{}
		return validator.check(varA, varB)
	case "max":
		validator := intMaxValidator{}
		return validator.check(varA, varB)
	case "min":
		validator := intMinValidator{}
		return validator.check(varA, varB)
	}
	return true
}

func validateStr(varA string, rule string, varB string) bool {
	switch rule {
	case "in":
		validator := stringInValidator{}
		return validator.check(varA, varB)
	case "len":
		validator := stringLenValidator{}
		return validator.check(varA, varB)
	case "regexp":
		validator := stringRegexpValidator{}
		return validator.check(varA, varB)
	}
	return true
}
