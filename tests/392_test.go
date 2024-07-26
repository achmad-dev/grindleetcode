package tests__test

import (
	"grindleetcode/problems"
	"reflect"
	"testing"
)

type testCase392 struct {
	S        string
	T        string
	Expected bool
}

func Test392(t *testing.T) {
	tests := []testCase392{
		{
			S:        "abc",
			T:        "ahbgdc",
			Expected: true,
		},
	}
	for _, tt := range tests {
		t.Run("solution", func(t *testing.T) {
			result := problems.IsSubsequence(tt.S, tt.T)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("IsSubsequence(%s, %s) = %v; want %v", tt.S, tt.T, result, tt.Expected)
			}
		})
	}
}
