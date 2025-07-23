package utils

func IsSome(arr []string, x string) bool {
	for i := range arr {
		if arr[i] == x {
			return true
		}
	}

	return false
}
