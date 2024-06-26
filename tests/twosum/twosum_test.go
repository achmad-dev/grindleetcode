package twosum_test

import (
	"grindleetcode/problems/twosum"
	"reflect"
	"testing"
)

type TestCase struct {
	Nums     []int
	Target   int
	Expected []int
}

func GenerateTestCase() []TestCase {
	testCases := []TestCase{
		{
			Nums:     []int{2, 7, 11, 15},
			Target:   9,
			Expected: []int{0, 1},
		},
		{
			Nums:     []int{3, 2, 4},
			Target:   6,
			Expected: []int{1, 2},
		},
		{
			Nums:     []int{3, 3},
			Target:   6,
			Expected: []int{0, 1},
		},
	}
	return testCases
}

func TestBruteForce(t *testing.T) {
	tests := GenerateTestCase()
	for _, tt := range tests {
		t.Run("brute force solution", func(t *testing.T) {
			result := twosum.BruteForce(tt.Nums, tt.Target)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("TwoSum(%d, %d) = %d; want %d", tt.Nums, tt.Target, result, tt.Expected)
			}
		})
	}
}

func TestOptimal(t *testing.T) {
	tests := GenerateTestCase()
	for _, tt := range tests {
		t.Run("optimal solution", func(t *testing.T) {
			result := twosum.Optimal(tt.Nums, tt.Target)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("TwoSum(%d, %d) = %d; want %d", tt.Nums, tt.Target, result, tt.Expected)
			}
		})
	}
}
