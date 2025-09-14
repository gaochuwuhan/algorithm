package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	roman := "III"
	res := RomanToInt(roman)
	assert.Equal(t, res, 3)
}
