package reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// CopyHelper helps to reflect copy (inject) some values
func CopyHelper(key string, from, to interface{}) error {
	// Checking reference type {to}
	typeOfTo := reflect.TypeOf(to)
	if typeOfTo.Kind() != reflect.Ptr {
		return fmt.Errorf(
			"Target must be pointer, but %T received",
			to,
		)
	}

	// Checking simple copy
	ok, err := simpleCopy(from, to)
	if err != nil {
		return fmt.Errorf("Unable to read %s. %s", key, err.Error())
	}
	if ok {
		return nil
	}

	// Checking string copy
	ok, err = smartStringCopy(from, to)
	if err != nil {
		return fmt.Errorf("Unable to read %s. %s", key, err.Error())
	}
	if ok {
		return nil
	}

	return fmt.Errorf("Unable to copy value of type %T to %T for key \"%s\"", from, to, key)
}

// simpleCopy copies values as-is if types are equal
func simpleCopy(from, to interface{}) (bool, error) {
	if reflect.TypeOf(from).Kind() == reflect.Ptr {
		return simpleCopy(reflect.ValueOf(from).Elem().Interface(), to)
	}
	if reflect.TypeOf(to) == reflect.PtrTo(reflect.TypeOf(from)) {
		reflect.ValueOf(to).Elem().Set(reflect.ValueOf(from))
		return true, nil
	}

	return false, nil
}

func smartStringCopy(from, to interface{}) (bool, error) {
	if reflect.TypeOf(from).Kind() == reflect.Ptr {
		return simpleCopy(reflect.ValueOf(from).Elem().Interface(), to)
	}
	if reflect.TypeOf(from).Kind() != reflect.String {
		return false, nil
	}

	sv, _ := from.(string)

	typeOfTo := reflect.TypeOf(to).Elem().Kind()
	valPoint := reflect.ValueOf(to).Elem()

	if typeOfTo == reflect.Bool {
		svl := strings.ToLower(sv)
		if svl == "true" || svl == "yes" || sv == "1" {
			valPoint.SetBool(true)
			return true, nil
		}
		if svl == "false" || svl == "no" || sv == "0" {
			valPoint.SetBool(false)
			return true, nil
		}

		return false, fmt.Errorf(
			"Unable to map string into boolean. Value \"%s\" is invalid",
			sv,
		)
	}

	if typeOfTo == reflect.Int {
		i, err := strconv.Atoi(sv)
		if err != nil {
			return false, fmt.Errorf(
				"Unable to map string to int. %s",
				err.Error(),
			)
		}

		valPoint.SetInt(int64(i))
		return true, nil
	}

	if typeOfTo == reflect.Int64 {
		i, err := strconv.ParseInt(sv, 10, 64)
		if err != nil {
			return false, fmt.Errorf(
				"Unable to map string to int64. %s",
				err.Error(),
			)
		}

		valPoint.SetInt(i)
		return true, nil
	}

	return false, nil
}
