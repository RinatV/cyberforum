package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Заменить все элементы, которые меньше последнего элемента на значение первого элемента. */
func TestTask(t *testing.T) {
	assert := assert.New(t)
	arr := [...]int{0, 1, 2, 3, 3}
	LessRepl(arr[:])
	assert.Equal([5]int{0, 0, 0, 3, 3},arr)
}
