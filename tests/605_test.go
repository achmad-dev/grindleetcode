package tests__test

import (
	"grindleetcode/problems"
	"reflect"
	"testing"
)

type testCase605 struct {
	Flowerbed []int
	N         int
	Expected  bool
}

func Test605(t *testing.T) {
	tests := []testCase605{
		{
			Flowerbed: []int{1, 0, 0, 0, 1},
			N:         1,
			Expected:  true,
		},
	}
	for _, tt := range tests {
		t.Run("solution", func(t *testing.T) {
			result := problems.CanPlaceFlowers(tt.Flowerbed, tt.N)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("CanPlaceFlowers(%v, %v) = %v; want %v", tt.Flowerbed, tt.N, result, tt.Expected)
			}
		})
	}
}
