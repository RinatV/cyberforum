package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	fill(arr[:])
	print(arr[:])
	BesideRepl(arr[:])
	print(arr[:])
}

func BesideRepl(arr []int) {
	mem := make([]int, len(arr))
	copy(mem, arr)
	var i, e int
	for i, e = range mem {
		if e == 0 {
			if i-1 >= 0 {
				arr[i-1] = 5
			}
			if i+1 < len(arr) {
				arr[i+1] = 5
			}
		}
	}
}

func fill(arr []int) {
	for i, _ := range arr {
		arr[i] = rand.Intn(9)
	}
}

func print(arr []int) {
	for i, _ := range arr {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println("")
}
