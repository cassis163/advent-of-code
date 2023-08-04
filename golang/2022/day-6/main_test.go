package main

import "testing"

func TestPartOne(t *testing.T) {
	testData := map[string]string{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "5",
		"nppdvjthqldpwncqszvftbrmjlhg":      "6",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "10",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "11",
	}

	for data, want := range testData {
		got := PartOne(data)
		if got != want {
			t.Errorf("PartOne(%s) = %s; want %s", data, got, want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	testData := map[string]string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    "19",
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "23",
		"nppdvjthqldpwncqszvftbrmjlhg":      "23",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "29",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "26",
	}

	for data, want := range testData {
		got := PartTwo(data)
		if got != want {
			t.Errorf("PartTwo(%s) = %s; want %s", data, got, want)
		}
	}
}
