package hw09structvalidator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var Builder strings.Builder
	for _, err := range v {
		Builder.WriteString(err.Err.Error())
	}
	return Builder.String()
}

type MinError struct{}

func (err MinError) Error() string {
	return "MinError"
}

type MaxError struct{}

func (err MaxError) Error() string {
	return "MaxError"
}

type LenError struct{}

func (err LenError) Error() string {
	return "LenError"
}

type InError struct{}

func (err InError) Error() string {
	return "InError"
}

type RegExpError struct{}

func (err RegExpError) Error() string {
	return "RegExpError"
}

type NotStructError struct{}

func (err NotStructError) Error() string {
	return "NotStructError"
}

type NotSupportedTypeError struct{}

func (err NotSupportedTypeError) Error() string {
	return "NotSupportedTypeError"
}

func Validate(v interface{}) error {
	var allErrors ValidationErrors
	reflectValue := reflect.ValueOf(v)
	if reflectType := reflectValue.Kind().String(); reflectType != "struct" {
		return NotStructError{}
	}
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Type().Field(i)
		if validate := field.Tag.Get("validate"); validate != "" {
			conditions := strings.Split(validate, "|")
			for _, condition := range conditions {
				splittedCondition := strings.Split(condition, ":")
				criteria := splittedCondition[0]
				criteriaValue := splittedCondition[1]
				switch reflectValue.Field(i).Interface().(type) {
				case int:
					allErrors = ValidateInt(criteria, field.Name, criteriaValue, reflectValue.Field(i).Int(), allErrors)
				case string:
					allErrors = ValidateString(criteria, field.Name, criteriaValue, reflectValue.Field(i).String(), allErrors)
				case []int:
					for _, value := range reflectValue.Field(i).Interface().([]int) {
						allErrors = ValidateInt(criteria, field.Name, criteriaValue, int64(value), allErrors)
					}
				case []string:
					for _, value := range reflectValue.Field(i).Interface().([]string) {
						allErrors = ValidateString(criteria, field.Name, criteriaValue, value, allErrors)
					}
				}
			}
		}
	}
	return allErrors
}

func ValidateInt(criteria, fieldName, criteriaValue string, toValidate int64,
	allErrors ValidationErrors,
) ValidationErrors {
	switch criteria {
	case "min":
		expected, _ := strconv.ParseInt(criteriaValue, 10, 64)
		if toValidate < expected {
			allErrors = append(allErrors, ValidationError{fieldName, MinError{}})
		}
	case "max":
		expected, _ := strconv.ParseInt(criteriaValue, 10, 64)
		if toValidate > expected {
			allErrors = append(allErrors, ValidationError{fieldName, MaxError{}})
		}
	case "in":
		setNumbers := strings.Split(criteriaValue, ",")
		flag := false
		for _, num := range setNumbers {
			expected, _ := strconv.ParseInt(num, 10, 64)
			if expected == toValidate {
				flag = true
			}
		}
		if !flag {
			allErrors = append(allErrors, ValidationError{fieldName, InError{}})
		}
	}
	return allErrors
}

func ValidateString(criteria, fieldName, criteriaValue, toValidate string,
	allErrors ValidationErrors,
) ValidationErrors {
	switch criteria {
	case "len":
		expected, _ := strconv.ParseInt(criteriaValue, 10, 64)
		if len(toValidate) != int(expected) {
			allErrors = append(allErrors, ValidationError{fieldName, LenError{}})
		}
	case "regexp":
		reg, _ := regexp.Compile(criteriaValue)
		if finded := reg.FindString(toValidate); finded == "" {
			allErrors = append(allErrors, ValidationError{fieldName, RegExpError{}})
		}
	case "in":
		setStrings := strings.Split(criteriaValue, ",")
		flag := false
		for _, str := range setStrings {
			if str == toValidate {
				flag = true
			}
		}
		if !flag {
			allErrors = append(allErrors, ValidationError{fieldName, InError{}})
		}
	}
	return allErrors
}
