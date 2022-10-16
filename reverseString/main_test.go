package binarySearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		l := len(s) - 1
		temp := s[l-i]
		s[l-i] = s[i]
		s[i] = temp
	}
}

func TestReverseString(t *testing.T) {
	cases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "1",
			input:    []byte{},
			expected: []byte{},
		},
	}

	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			reverseString(cc.input)
			require.Equal(t, cc.expected, cc.input)
		})
	}
}
