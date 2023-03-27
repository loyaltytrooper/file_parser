package helper

func GetDescription(arr []string) string {
	if len(arr) == 4 {
		return arr[1]
	} else if len(arr) == 5 {
		return arr[1] + arr[2]
	} else if len(arr) == 6 {
		return arr[1] + arr[2] + arr[3]
	} else {
		return ""
	}
}
