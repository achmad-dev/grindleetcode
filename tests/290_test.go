package tests__test

import (
	"grindleetcode/problems"
	"testing"
)

func Test290(t *testing.T) {
	result := problems.WordPattern("abba", "dog cat cat dog")
	if !result {
		t.Errorf("WordPattern() = %v; want %v", result, true)
	}
}
