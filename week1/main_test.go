package main

import "testing"

func TestSum(t *testing.T) {
	list, err := ReturnIntList("1,2,4,5,6")
	if err != nil {
		t.Fail()
	}
	if len(list) != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", len(list), 5)
	}
}
