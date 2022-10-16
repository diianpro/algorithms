package longestSubstr

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	if s == "" {
		return 0
	}

	start := 0
	end := 1
	maxLen := start
	for end < len(s) {
		sub := s[start:end]
		ss := string(s[end])
		if !strings.Contains(sub, ss) {
			end++
			if end-start > maxLen {
				maxLen = end - start
			}
		} else {
			start += strings.Index(sub, ss) + 1
		}
	}
	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {

	cases := []struct {
		input    string
		expected int
	}{
		{
			input:    "aleiivuuxszpaqojv",
			expected: 10,
		},
		{
			input:    "ckilbkd",
			expected: 5,
		},
		{
			input:    "abba",
			expected: 2,
		},
		{
			input:    "dvdf",
			expected: 3,
		},
		{
			input:    "dsvdf",
			expected: 4,
		},
		{
			input:    " ",
			expected: 1,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
		{
			input:    "pwwkwe",
			expected: 3,
		},
	}

	for _, cc := range cases {
		tt := t
		tt.Run(cc.input, func(t *testing.T) {
			shortestSubStrLen := lengthOfLongestSubstring(cc.input)
			require.Equal(t, cc.expected, shortestSubStrLen)
		})
	}
}
