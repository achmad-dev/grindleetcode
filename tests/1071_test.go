package tests__test

import (
	"grindleetcode/problems"
	"reflect"
	"testing"
)

type TestCase1071 struct {
	Word1    string
	Word2    string
	Expected string
}

func GenerateTestCase1071() []TestCase {
	testCase := []TestCase{
		{
			Word1:    "ABCABC",
			Word2:    "ABC",
			Expected: "ABC",
		},
	}
	return testCase
}

func Test1071(t *testing.T) {
	tests := GenerateTestCase1071()
	for _, tt := range tests {
		t.Run("solution", func(t *testing.T) {
			result := problems.GcdOfStrings(tt.Word1, tt.Word2)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("GcdOfStrings(%s, %s) = %s; want %s", tt.Word1, tt.Word2, result, tt.Expected)
			}
		})
	}
}
