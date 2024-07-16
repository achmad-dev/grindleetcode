package movezero_test

import (
	movezero "grindleetcode/problems/move_zero"
	"reflect"
	"testing"
)

type TestCase struct {
	Nums     []int
	Expected []int
}

func GenerateTestCase() []TestCase {
	testCases := []TestCase{
		{
			Nums:     []int{1, 0, 2, 3, 0, -4},
			Expected: []int{1, 2, 3, -4, 0, 0},
		},
	}
	return testCases
}

func TestBruteForce(t *testing.T) {
	tests := GenerateTestCase()
	for _, tt := range tests {
		t.Run("brute force solution", func(t *testing.T) {
			result := movezero.MoveZeroBruteForce(tt.Nums)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("MoveZero(%d) = %d; want %d", tt.Nums, result, tt.Expected)
			}
		})
	}
}
