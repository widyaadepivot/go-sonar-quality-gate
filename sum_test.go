package gosonarqualitygate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		input      []int
		wantResult int
	}{
		{
			input:      nil,
			wantResult: 0,
		},
		{
			input:      []int{},
			wantResult: 0,
		},
		{
			input:      []int{1, 2, 3},
			wantResult: 6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.wantResult, Sum(tt.input...))
	}
}
