package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [10] int
	rand.Seed(time.Now().UnixNano())
	fill(arr[:])
	print(arr[:])
	DublToZero(arr[:])
	print(arr[:])
}

func DublToZero(arr []int) {
	var i, j, c int
	for i = 0; i < len(arr); i++ {
		c = 0
		for j = i; j < len(arr); j++ {
			if arr[i] == arr[j] {
				c++
			}
		}
		if c >= 2 {
			for j = i + 1; j < len(arr); j++ {
				if arr[i] == arr[j] {
					arr[j] = 0
				}
			}
			arr[i] = 0
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
