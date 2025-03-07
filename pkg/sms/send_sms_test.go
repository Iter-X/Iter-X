package sms_test

import (
	"testing"

	"github.com/iter-x/iter-x/pkg/sms"
)

func TestGenerateRandomNumberCode(t *testing.T) {
	tests := []struct {
		length int
		want   int
	}{
		{5, 5},
		{10, 10},
		{0, 0},
		{-1, 0},
	}

	for _, test := range tests {
		result := sms.GenerateRandomNumberCode(test.length)
		if len(result) != test.want {
			t.Errorf("GenerateRandomNumberCode(%d) = %s; want length %d", test.length, result, test.length)
		}
	}
}
