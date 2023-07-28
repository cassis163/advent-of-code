package util

import "os"

func ReadFileAsString(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(file), nil
}
