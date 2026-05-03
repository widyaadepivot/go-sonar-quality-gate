package gosonarqualitygate_test

import (
	"testing"

	. "go-sonar-quality-gate"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	tests := []struct {
		values     []int
		wantMax    int
		wantResult bool
	}{
		{
			values:     nil,
			wantMax:    0,
			wantResult: false,
		},
		{
			values:     []int{},
			wantMax:    0,
			wantResult: false,
		},
		{
			values:     []int{1, 3, 2},
			wantMax:    3,
			wantResult: true,
		},
	}
	for _, tt := range tests {
		max, ok := Max(tt.values...)
		assert.Equal(t, tt.wantMax, max)
		assert.Equal(t, tt.wantResult, ok)
	}
}
