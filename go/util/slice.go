package util

import "errors"

func GetMaxIntInSlice(ints []int) (int, error) {
	if len(ints) == 0 {
		return 0, errors.New("slice is empty")
	}

	var max int
	for _, i := range ints {
		if i > max {
			max = i
		}
	}

	return max, nil
}

func GetIndexInSlice(slice []string, value string) (int, error) {
	for i, v := range slice {
		if v == value {
			return i, nil
		}
	}

	return 0, errors.New("value not found")
}

func IsValueInSlice(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
