package tests__test

import (
	"grindleetcode/problems"
	"reflect"
	"testing"
)

type SockPairsTestCase struct {
	Ar       []int32
	N        int32
	Expected int32
}

func TestSockMerchant(t *testing.T) {
	tests := []SockPairsTestCase{
		{
			N:        9,
			Ar:       []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			Expected: 3,
		},
		{
			N:        7,
			Ar:       []int32{1, 2, 1, 2, 1, 3, 2},
			Expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run("solution", func(t *testing.T) {
			result := problems.SockMerchant(tt.N, tt.Ar)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("SockMerchant(%d, %d) = %d; want %d", tt.N, tt.Ar, result, tt.Expected)
			}
		})
	}
}
