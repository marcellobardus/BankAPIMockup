package utils

func IsInArray(element string, array []string) bool {

	if len(array) == 0 {
		return false
	}

	for _, i := range array {
		if i == element {
			return true
		}
	}

	return false
}
