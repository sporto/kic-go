package matchers

import (
	"github.com/onsi/gomega"
  "time"
  // "reflect"
)


func BeWithin(expected time.Time) gomega.OmegaMatcher {
	return &withinMatcher{
		expected: expected,
	}
}

type withinMatcher struct {
	expected time.Time
}

func (matcher *withinMatcher) Match(actual interface{}) (success bool, message string, err error) {
  actual2 := actual.(time.Time)
  dif := actual2.Sub(matcher.expected)
  if dif > 100 {
    return false, "Duration not within range", nil
  }
	return true, "", nil
}
