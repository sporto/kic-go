package misc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateInterest(t *testing.T) {
	i := CalculateInterest(100, 365, 3)
	assert.Equal(t, i, 3, "they should be equal")
}

func TestCalculateInterest2(t *testing.T) {
	i := CalculateInterest(100, 10, 3)
	assert.Equal(t, i, 0.08219178, "they should be equal")
}
