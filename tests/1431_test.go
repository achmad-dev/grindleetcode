package tests__test

import (
	"grindleetcode/problems"
	"reflect"
	"testing"
)

type testCase1431 struct {
	candies  []int
	extra    int
	expected []bool
}

func Test1431(t *testing.T) {
	tests := []testCase1431{
		{
			candies:  []int{2, 3, 5, 1, 3},
			extra:    3,
			expected: []bool{true, true, true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run("solution", func(t *testing.T) {
			result := problems.KidsWithCandies(tt.candies, tt.extra)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KidsWithCandies(%d, %d) = %v; want %v", tt.candies, tt.extra, result, tt.expected)
			}
		})
	}
}
