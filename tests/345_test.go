package tests__test

import (
	"grindleetcode/problems"
	"reflect"
	"testing"
)

type testCase345 struct {
	Input    string
	Expected string
}

func Test345(t *testing.T) {
	tests := []testCase345{
		{
			Input:    "hello",
			Expected: "holle",
		},
	}
	for _, tt := range tests {
		t.Run("solution", func(t *testing.T) {
			result := problems.ReverseVowels(tt.Input)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("ReverseVowels(%s) = %s; want %s", tt.Input, result, tt.Expected)
			}
		})
	}
}
