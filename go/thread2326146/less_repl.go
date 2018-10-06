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
	LessRepl(arr[:])
	print(arr[:])
}

func LessRepl(arr []int) {
	var i int
	for i = 1; i < len(arr)-1; i++ {
		if arr[i] < arr[len(arr)-1] {
			arr[i] = arr[0]
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
