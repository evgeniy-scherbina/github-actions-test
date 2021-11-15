package main

import "testing"

func TestSum(t *testing.T) {
	if sum(2, 3) != 5 {
		t.FailNow()
	}
}
