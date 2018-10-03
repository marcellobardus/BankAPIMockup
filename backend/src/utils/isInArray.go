package utils

import (
	"reflect"
)

func isInArray(element interface{}, array []interface{}) bool {

	if len(array) == 0 {
		return false
	}

	if reflect.TypeOf(element) != reflect.TypeOf(array[0]) {
		panic("Incompatible types")
	}

	for _, i := range array {
		if i == element {
			return true
		}
	}

	return false
}
