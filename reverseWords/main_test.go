package binarySearch

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	res := make([]string, 0, len(ss))
	for _, str := range ss {
		bstr := []byte(str)
		for i := 0; i < len(bstr)/2; i++ {
			l := len(bstr) - 1
			temp := bstr[l-i]
			bstr[l-i] = bstr[i]
			bstr[i] = temp
		}
		res = append(res, string(bstr))
	}
	return strings.Join(res, " ")
}

func TestReverseWords(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "1",
			input:    "ab's cde",
			expected: "s'ba edc",
		},
	}

	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.expected, reverseWords(cc.input))
		})
	}
}
