package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Обнулить элементы массива, которые повторяются больше 2-х раз. (1задание) */
func TestTask(t *testing.T) {
	assert := assert.New(t)
	arr := [...]int{0, 1, 2, 3, 3}
	DublToZero(arr[:])
	assert.Equal([5]int{0, 1, 2, 0, 0},arr)
}
