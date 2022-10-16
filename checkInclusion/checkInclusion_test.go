package longestSubstr

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func checkInclusion(s1 string, s2 string) bool {
	start := 0
	end := len(s1) - 1
	contains := true
	for end < len(s2) {
		sub := s2[start:end]
		for _, s := range s1 {
			contains = contains && strings.Contains(sub, string(s))
		}
		if contains {
			return true
		}
		start++
		end++
	}
	return false
}

func TestLengthOfLongestSubstring(t *testing.T) {

	cases := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "1",
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
	}

	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.expected, checkInclusion(cc.s1, cc.s2))
		})
	}
}
