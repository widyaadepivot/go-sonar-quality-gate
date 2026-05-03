package gosonarqualitygate_test

import (
	. "go-sonar-quality-gate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		input      []string
		wantResult string
	}{
		{
			input:      nil,
			wantResult: "",
		},
		{
			input:      []string{},
			wantResult: "",
		},
		{
			input:      []string{"Hello", "", ", world!"},
			wantResult: "Hello, world!",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.wantResult, Concat(tt.input...))
	}
}
