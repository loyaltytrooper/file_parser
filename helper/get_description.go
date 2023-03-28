package helper

import "strings"

func GetDescription(arr []string) []string {
	var description string
	if len(arr) == 4 {
		description = arr[1]
	} else if len(arr) == 5 {
		description = arr[1] + arr[2]
	} else if len(arr) == 6 {
		description = arr[1] + arr[2] + arr[3]
	} else {
		description = ""
	}

	return strings.Split(description, "/")
}
