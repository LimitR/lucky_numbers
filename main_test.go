package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxNumOne(t *testing.T) {
	num := getMaxNum(10, 3)
	assert.Equal(t, num, 334, "they should be equal")
	assert.True(t, checkValid(10, 334), "True is true!")
}

func TestGetMaxNumTwo(t *testing.T) {
	num := getMaxNum(27, 3)
	assert.Equal(t, num, 999, "they should be equal")
	assert.True(t, checkValid(27, 999), "True is true!")
}

func TestGetMaxNumThree(t *testing.T) {
	num := getMaxNum(32, 5)
	assert.Equal(t, num, 66677, "they should be equal")
	assert.True(t, checkValid(32, 66677), "True is true!")
}
