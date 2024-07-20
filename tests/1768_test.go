package tests__test

import (
	"grindleetcode/problems"
	"reflect"
	"testing"
)

type TestCase struct {
	Word1    string
	Word2    string
	Expected string
}

func GenerateTestCase() []TestCase {
	testCase := []TestCase{
		{
			Word1:    "abc",
			Word2:    "pqr",
			Expected: "apbqcr",
		},
	}
	return testCase
}

func Test1768(t *testing.T) {
	tests := GenerateTestCase()
	for _, tt := range tests {
		t.Run("solution", func(t *testing.T) {
			result := problems.MergeAlternately(tt.Word1, tt.Word2)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("MergeAlternately(%s, %s) = %s; want %s", tt.Word1, tt.Word2, result, tt.Expected)
			}
		})
	}
}
