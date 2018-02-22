package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHelper(t *testing.T) {
	tests := []struct {
		item   []string
		result Helper
		err    error
	}{
		{item: []string{""},
			result: Helper{},
			err:    nil},
	}
	for _, test := range tests {
		result, err := newHelper(test.item[0])
		assert.IsType(t, test.err, err)
		assert.IsType(t, &test.result, result)
	}
}
