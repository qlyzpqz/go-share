package main

import (
	"testing"
)

func TestIsNegative(t *testing.T) {
	if is_negative(-1) == false {
		t.Error("is_negative return failed")
	}

	if is_negative(1) == true {
		t.Error("is_negative return failed")
	}
}

func BenchmarkIsNegative(t *testing.B) {
	for i := 0; i < t.N; i++ {
		is_negative(-1)
	}
}
