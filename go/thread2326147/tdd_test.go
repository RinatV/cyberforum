package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Найти все нулевые элементы и заменить значение соседних 2-х элементов на 5. */
func TestTask(t *testing.T) {
	assert := assert.New(t)
	arr := [...]int{0, 1, 0, 0, 3}
	BesideRepl(arr[:])
	//                 {0, 1, 0, 0, 3}
	assert.Equal([5]int{0, 5, 5, 5, 5}, arr)
}
