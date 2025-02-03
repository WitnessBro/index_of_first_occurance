package indexoffirstoccurance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var haystack = "sabutsa"
	var needle = "sad"
	var k = strStr(haystack, needle)

	assert.Equal(t, k, 3)
}
func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	lps := compute_lps(needle)
	i := 0
	j := 0

	for i < n {
		if haystack[i] == needle[j] {
			i += 1
			j += 1
			if j == m {
				return i - j
			}
		} else {
			if j > 0 {
				j = lps[j-1]
			} else {
				i += 1
			}
		}
	}
	return -1
}

func compute_lps(pattern string) []int {
	pattern_length := len(pattern)
	lps := make([]int, pattern_length)
	j := 0
	i := 1

	for i < pattern_length {
		if pattern[i] == pattern[j] {
			j += 1
			lps[i] = j
			i += 1
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i += 1
			}
		}
	}
	return lps
}
