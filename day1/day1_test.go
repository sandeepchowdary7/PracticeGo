package day1

import (
	"testing"
)

func TestInitDay1(t *testing.T) {
	t.Run("Negative_Case", func(t *testing.T) {
		input1 := []int{1, 2, 3, 5, 6, 9}
		if a := checkSum(input1, 20); a {
			t.Error("Test Negative on Positive Case")
		}
	})

	t.Run("Positive_Case", func(t *testing.T) {
		input2 := []int{1, 2, 7, 5, 6, 13}
		if a := checkSum(input2, 20); a {
			t.Log("Test Passed on Positive Case")
		}
	})
}
