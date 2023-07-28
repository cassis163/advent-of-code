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
