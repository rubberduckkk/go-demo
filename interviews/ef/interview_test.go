package ef

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestNonrepeatedSubstrings(t *testing.T) {
	type testCase struct {
		s           string
		expectedAns []string
	}

	cases := []testCase{
		{
			s:           "abacb",
			expectedAns: []string{"bac", "acb"},
		},
		{
			s:           "accbddcb",
			expectedAns: []string{"cbd", "dcb"},
		},
	}

	for _, c := range cases {
		ans := longestNonrepeatedSubstrings(c.s)
		assert.Equal(t, c.expectedAns, ans)
	}
}
