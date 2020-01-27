package main

import "testing"

const (
	success = "\u2713"
	failure = "\u2717"
)
func TestCountDigits(t *testing.T) {

	testData := map[string]struct { in, expected int} {
		"one": {1,1},
		"nine": {9,1},
		"ten": {10,2},
		"99": {99,2},
		"billion": {1e9,10},
		"almost billion": {1e9-1,9},
	}

	t.Log("Given the need to test counting of digits...")
	for name, data := range testData {
		//t.Log("\tTest %q\tchecking that %d has %d digits", name, data.in, data.out)
		got := countDigits(data.in)
		if got != data.expected {
			t.Errorf("%s %-20q For %d expected %d but got %d", failure, name, data.in, data.expected, got)
		} else {
			t.Logf("%s %-20q For %d got expected value %d", success, name, data.in, got)
		}
	}
}
