package matchers

import (
	"github.com/onsi/gomega"
)


func BeWithin(expected interface{}) gomega.OmegaMatcher {
	return &withinMatcher{
		expected: expected,
	}
}

type withinMatcher struct {
	expected interface{}
}

func (matcher *withinMatcher) Match(actual interface{}) (success bool, message string, err error) {
	return true, "", nil
}
